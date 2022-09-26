package mysql

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type DB struct {
	*sql.DB
	sqlExecutionTimeout int
}

type DbMeta struct {
	User              string
	Password          string
	Socket            string
	Host              string
	Port              string
	Version           string
	ConnectTimeout    int
	ExecSqlTimeout    int
	ExtraSysVariables string
}

func PingDbOrTimeout(db *DB, timeoutSeconds int) error {
	return pingDbOrTimeout(db.DB, timeoutSeconds)
}

func pingDbOrTimeout(db *sql.DB, timeoutSeconds int) error {
	var pingErr error
	fin := <-util.Timeout(func() {
		pingErr = util.Ping(db)
	}, timeoutSeconds)
	if !fin {
		return errors.New("ping db timeout")
	}
	return pingErr
}

// OpenDb如果version字段未传递,则不会指定session的MaxStatementTime配置
func OpenDb(stage *log.Stage, user, pass, socket, host, port string, connectTimeout, execSqlTimeout int, version ...string) (db *DB, err error) {
	var sysVariables = new(SysVariables)
	if len(version) > 0 {
		mysqlVersion, err := util.NewVersion(version[0])
		if err != nil {
			return nil, err
		}
		*sysVariables = sysVariables.SetMaxStatementTime(mysqlVersion)
	}
	key := DbMeta{
		User:              user,
		Password:          pass,
		Socket:            socket,
		Host:              host,
		Port:              port,
		ConnectTimeout:    connectTimeout,
		ExecSqlTimeout:    execSqlTimeout,
		ExtraSysVariables: sysVariables.String(),
	}
	return OpenDbWithDbMeta(stage, key)
}

/* workflow of OpenDbWithDbMeta, A B request for same db connection
              +                  +
            req A              req B
              |                  |
              |                  |
+------+------v---------+-----------------+---------------+
       |                |        |        |
       |                |        |      check
       |                |        |     old conn
       |             dial and    v        |
     check            check               v
    old conn         new conn             |
       |                |                 |
       |      or        |       or    dial and
       |                |              check
       |                |             new conn
       |                |                 |
       |                |                 |
+------v------+---------v--------+--------v----------------+
              |                  |
              A                  B
              v                  v
*/
var dbCache = make(map[DbMeta]*DB)
var DbCacheMutex = sync.Mutex{}
var healthCheckers = make(map[DbMeta]chan struct{})
var healthCheckersMutex = sync.Mutex{}
var once = sync.Once{}
var sqlExecTimeoutError = errors.New("sql-execution-timeout")

type SysVariables []string

func (s SysVariables) SetReplicationConsistency(currentVersion util.Version, currentStandard, uguardGR string) SysVariables {
	if currentStandard == uguardGR {
		introducedVersion := util.NewVersionWithNumbers(8, 0, 14)
		if currentVersion.NoLessThan(introducedVersion) {
			s = append(s, "group_replication_consistency=EVENTUAL")
		}
	}
	return s
}

func (s SysVariables) SetMaxStatementTime(currentVersion util.Version) SysVariables {
	// 不限制mysql语句执行时间,这里只限制了max_execution_time,如果语句执行时间超过timeout其实还是会超时的
	// 此处设置是为了防止实例上配置该参数过小导致解析数据集出现panic
	version574 := util.NewVersionWithNumbers(5, 7, 4)
	version577 := util.NewVersionWithNumbers(5, 7, 7)
	version578 := util.NewVersionWithNumbers(5, 7, 8)
	if currentVersion.IsInTheRange(version574, version577) {
		s = append(s, "max_statement_time=0")
	}
	heightVersion := version578
	if currentVersion.NoLessThan(heightVersion) {
		s = append(s, "max_execution_time=0")
	}
	return s
}

func (s SysVariables) String() string {
	return strings.Join(s, "&")
}

func OpenDbWithDbMeta(stage *log.Stage, key DbMeta) (db *DB, err error) {
	once.Do(func() {
		go chCacheRotate()
		go cleanupDB()
	})

	ch := startCheckDbConn(stage, key)
	<-ch

	DbCacheMutex.Lock()
	if db, ok := dbCache[key]; ok {
		DbCacheMutex.Unlock()
		return db, nil
	}
	DbCacheMutex.Unlock()

	return nil, errors.New("no pingable connection and dial new connection fail")
}

func startCheckDbConn(stage *log.Stage, key DbMeta) <-chan struct{} {
	healthCheckersMutex.Lock()
	if finishCh, ok := healthCheckers[key]; ok {
		healthCheckersMutex.Unlock()
		return finishCh
	}

	finishCh := make(chan struct{})
	healthCheckers[key] = finishCh
	healthCheckersMutex.Unlock()

	go checkDbConn(stage, key, finishCh)
	return finishCh
}

func checkDbConn(stage *log.Stage, key DbMeta, finishCh chan struct{}) {
	defer func() {
		healthCheckersMutex.Lock()
		delete(healthCheckers, key)
		healthCheckersMutex.Unlock()
		close(finishCh)
	}()

	for {
		DbCacheMutex.Lock()
		_, cacheFound := dbCache[key]

		// clean unclean cache, only a connection use old password
		if !cacheFound {
			for existKey := range dbCache {
				if key.Socket == existKey.Socket && key.Host == existKey.Host &&
					key.Port == existKey.Port && key.User == existKey.User &&
					key.ExecSqlTimeout == existKey.ExecSqlTimeout &&
					key.ConnectTimeout == existKey.ConnectTimeout &&
					key.ExtraSysVariables == existKey.ExtraSysVariables {

					dbCache[existKey].Close()
					log.Detail(stage, "db cache: [%+v] meta changed to [%+v] because of password changed.", existKey, key)
					delete(dbCache, existKey)
				}
			}
		}
		DbCacheMutex.Unlock()

		if cacheFound {
			return
		}

		// dial a new connection
		var newDb *DB
		var newDbErr error

		switch {
		case key.Socket != "" && key.ExtraSysVariables != "":
			newDb, newDbErr = OpenDbBySocketWithoutCacheWithExtraSysVars(stage, key.User, key.Password, key.Socket,
				key.ExtraSysVariables, key.ConnectTimeout, key.ExecSqlTimeout)

		case key.Socket != "" && key.ExtraSysVariables == "":
			newDb, newDbErr = OpenDbBySocketWithoutCache(stage, key.User, key.Password, key.Socket,
				key.ConnectTimeout, key.ExecSqlTimeout)

		case key.Socket == "" && key.ExtraSysVariables != "":
			newDb, newDbErr = OpenDbWithoutCacheWithExtraSysVars(stage, key.User, key.Password, key.Host, key.Port,
				key.ExtraSysVariables, key.ConnectTimeout, key.ExecSqlTimeout)

		case key.Socket == "" && key.ExtraSysVariables == "":
			newDb, newDbErr = OpenDbWithoutCache(stage, key.User, key.Password, key.Host, key.Port,
				key.ConnectTimeout, key.ExecSqlTimeout)
		}

		if nil != newDbErr {
			log.KeyDilute2(stage, "dial_"+key.Host+":"+key.Port,
				"dial %s:%s with %s fail: %v", key.Host, key.Port, key.User, newDbErr)
			return
		}

		newDb.SetMaxIdleConns(5)
		newDb.SetMaxOpenConns(20)

		DbCacheMutex.Lock()
		if _, foundAnotherCache := dbCache[key]; foundAnotherCache {
			newDb.Close()
			DbCacheMutex.Unlock()
			continue
		} else {
			dbCache[key] = newDb
			DbCacheMutex.Unlock()
			return
		}
	}
}

// for hash map do not shrink after item deleted, https://github.com/golang/go/issues/20135
func chCacheRotate() {
	for range time.Tick(time.Hour) {
		newCheckers := make(map[DbMeta]chan struct{})

		healthCheckersMutex.Lock()
		for k, v := range healthCheckers {
			newCheckers[k] = v
		}
		healthCheckers = newCheckers
		healthCheckersMutex.Unlock()
	}
}

// DMP-10864: clean db when instance is removed
func cleanupDB() {
	stage := log.NewStage().Enter("cleanupDB")
	for range time.Tick(time.Hour) {
		DbCacheMutex.Lock()
		for k, db := range dbCache {
			if nil != PingDbOrTimeout(db, 5) {
				db.Close()
				log.Detail(stage, "db cache: [%+v] meta is closed because of unhealthy.", k)
				delete(dbCache, k)
			}
		}
		DbCacheMutex.Unlock()
	}
}

func OpenDbWithoutCache(stage *log.Stage, user, pass, host, port string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	return openDbWithoutCache(stage, user, pass, host, port, "", "", connectTimeout, execHaSqlTimeout)
}

func OpenDbWithoutCacheWithSchema(stage *log.Stage, user, pass, host, port, schema string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	return openDbWithoutCache(stage, user, pass, host, port, schema, "", connectTimeout, execHaSqlTimeout)
}

func OpenDbWithoutCacheWithExtraSysVars(stage *log.Stage, user, pass, host, port, extraSysVariables string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	return openDbWithoutCache(stage, user, pass, host, port, "", extraSysVariables, connectTimeout, execHaSqlTimeout)
}

func openDbWithoutCache(stage *log.Stage, user, pass, host, port, schema, extraSysVariables string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	log.Detail(stage, "host=%v; port=%v; schema=%v; connectTimeout=%v; execHaSqlTimeout=%v", host, port, schema, connectTimeout, execHaSqlTimeout)
	userPass := user
	if "" != pass {
		userPass += ":" + pass
	}
	var connectCmd string
	if extraSysVariables != "" {
		connectCmd = fmt.Sprintf("%s@(%s:%s)/%s?timeout=%ds&autocommit=1&multiStatements=true&loc=Local&%s",
			userPass, host, port, schema, connectTimeout, extraSysVariables)
	} else {
		connectCmd = fmt.Sprintf("%s@(%s:%s)/%s?timeout=%ds&autocommit=1&multiStatements=true&loc=Local",
			userPass, host, port, schema, connectTimeout)
	}
	return openDbByConnStrWithoutCache(stage, connectCmd, host+":"+port, execHaSqlTimeout)
}

func OpenDbBySocketWithoutCache(stage *log.Stage, user, pass, socket string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	return openDbBySocketWithoutCache(stage, user, pass, socket, "", "", connectTimeout, execHaSqlTimeout)
}

func OpenDbBySocketWithoutCacheWithSchema(stage *log.Stage, user, pass, socket, schema string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	return openDbBySocketWithoutCache(stage, user, pass, socket, schema, "", connectTimeout, execHaSqlTimeout)
}

func OpenDbBySocketWithoutCacheWithExtraSysVars(stage *log.Stage, user, pass, socket, extraSysVariables string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	return openDbBySocketWithoutCache(stage, user, pass, socket, "", extraSysVariables, connectTimeout, execHaSqlTimeout)
}

func openDbBySocketWithoutCache(stage *log.Stage, user, pass, socket, schema, extraSysVariables string, connectTimeout, execHaSqlTimeout int) (db *DB, err error) {
	log.Detail(stage, "socket=%v; schema=%v; connectTimeout=%v; execHaSqlTimeout=%v", socket, schema, connectTimeout, execHaSqlTimeout)
	userPass := user
	if "" != pass {
		userPass += ":" + pass
	}

	var connectCmd string
	if extraSysVariables != "" {
		connectCmd = fmt.Sprintf("%s@unix(%s)/%s?timeout=%ds&autocommit=1&multiStatements=true&loc=Local&%s",
			userPass, socket, schema, connectTimeout, extraSysVariables)
	} else {
		connectCmd = fmt.Sprintf("%s@unix(%s)/%s?timeout=%ds&autocommit=1&multiStatements=true&loc=Local",
			userPass, socket, schema, connectTimeout)
	}
	return openDbByConnStrWithoutCache(stage, connectCmd, socket, execHaSqlTimeout)
}

func openDbByConnStrWithoutCache(stage *log.Stage, connectCmd, flag string, execHaSqlTimeout int) (db *DB, err error) {
	stage.Enter("SqlOpen")
	defer stage.Exit()

	started := time.Now()
	defer func() {
		if delta := time.Now().Sub(started).Seconds(); delta > 5 {
			log.Brief(stage, "!!! open db %v exceed 5s (actual %vs)", flag, int(delta))
		}
	}()

	sqlDb, err := sql.Open("mysql", connectCmd)
	if nil != err {
		return nil, err
	}
	if err := pingDbOrTimeout(sqlDb, 5 /* hard-coding */); nil != err {
		sqlDb.Close()
		return nil, err
	}
	db = &DB{sqlDb, execHaSqlTimeout}
	return db, nil
}

// always record input/output logs
func SqlQuery(stage *log.Stage, db *DB, s string, arguments ...interface{}) ([]map[string]string, error) {
	return SqlQueryWithLogLevelSpecified(stage, false, db, s, arguments...)
}

func SqlQueryWithLogLevelSpecified(stage *log.Stage, logOnErrOnly bool, db *DB, s string, arguments ...interface{}) ([]map[string]string, error) {
	stage.Enter("SqlQuery")
	defer stage.Exit()
	result := []map[string]string{}
	mutex := sync.Mutex{} //TRY to fix mystery
	err := sqlQueryForEachRow(stage, logOnErrOnly, db, s, []string{}, func(row map[string]string, hasNonEmptyField bool) bool {
		mutex.Lock()
		result = append(result, row)
		mutex.Unlock()
		return true
	}, arguments...)
	if nil != err {
		return nil, err
	}
	mutex.Lock()
	defer mutex.Unlock()
	return result, nil
}

// always record input/output logs
func SqlExec(stage *log.Stage, db *DB, sql string, args ...interface{}) (err error) {
	return SqlExecWithLogLevelSpecified(stage, false, db, sql, args...)
}

func SqlExecWithLogLevelSpecified(stage *log.Stage, logOnErrOnly bool, db *DB, sql string, args ...interface{}) (err error) {
	stage.Enter("SqlExec")
	defer stage.Exit()

	mutex := sync.Mutex{} //TRY to fix mystery
	log.Detail(stage, "sql[%v]args[%v]", sql, args)
	if !<-util.Timeout(func() {
		r, e := db.Exec(sql, args...)
		if r != nil && !logOnErrOnly {
			rowsAffected, _ := r.RowsAffected() // error occurs only when driver doesn't support this function, just ignore
			log.Detail(stage, "sql[%v]args[%v]executed,RowsAffected[%v]", sql, args, rowsAffected)
		}
		mutex.Lock()
		err = e
		mutex.Unlock()
	}, db.sqlExecutionTimeout) {
		log.Detail(stage, "sql[%v]args[%v]err:%v", sql, args, sqlExecTimeoutError)
		mutex.Lock()
		defer mutex.Unlock()
		return sqlExecTimeoutError
	} else {
		if err != nil {
			log.Detail(stage, "sql[%v]args[%v]err:%v", sql, args, err.Error())
		}
		mutex.Lock()
		defer mutex.Unlock()
		return err
	}
}

func SqlQueryRow(stage *log.Stage, db *DB, s string, arguments ...interface{}) (result map[string]string, err error) {
	return SqlQueryRow2(stage, db, s, []string{}, arguments...)
}

func SqlQueryRow2(stage *log.Stage, db *DB, s string, columns []string, arguments ...interface{}) (map[string]string, error) {
	var result map[string]string
	mutex := sync.Mutex{} //TRY to fix mystery
	err := sqlQueryForEachRow(stage, false, db, s, columns, func(row map[string]string, hasNonEmptyField bool) bool {
		mutex.Lock()
		if !hasNonEmptyField {
			result = nil
		} else {
			result = row
		}
		mutex.Unlock()
		return false
	}, arguments...)
	if nil != err {
		return nil, err
	}
	mutex.Lock()
	defer mutex.Unlock()
	return result, nil
}

func sqlQueryForEachRow(stage *log.Stage, logOnErrOnly bool, db *DB, s string, columns []string, iter func(row map[string]string, hasNonEmptyField bool) bool, arguments ...interface{}) (err error) {
	log.Detail(stage, "query[%v]args[%v]", s, arguments)
	rowsChan := make(chan *sql.Rows, 1)
	errChan := make(chan error, 2)

	if !<-util.Timeout(func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- errors.New("panic in sql.query")
				return
			}
		}()

		rows, err := db.Query(s, arguments...)
		if nil != err {
			errChan <- err
			return
		}
		rowsChan <- rows
	}, db.sqlExecutionTimeout) {
		log.Detail(stage, "query[%v]args[%v]err:%v", s, arguments, sqlExecTimeoutError)

		//finalizer
		go func() {
			select {
			case <-errChan:
			case rows := <-rowsChan:
				rows.Close()
			}
		}()

		return sqlExecTimeoutError
	}

	var rows *sql.Rows
	select {
	case err := <-errChan:
		log.Detail(stage, "query[%v]args[%v]err:%v", s, arguments, err)
		return err
	case rows = <-rowsChan:
	}

	sb := &strings.Builder{}
	fmt.Fprintf(sb, "query[%v]args[%v]result:", s, arguments)
	specLen := sb.Len() + 1000
	if nil != rows {
		defer rows.Close()
		cols, _ := rows.Columns()
		buf := make([]interface{}, len(cols))
		data := make([]sql.NullString, len(cols))
		for i := range buf {
			buf[i] = &data[i]
		}

		ri := 0
		for rows.Next() {
			rows.Scan(buf...)
			fprintfWhileLenLessThanSpecifiedInt(sb, specLen, "[%v]", ri)
			row := make(map[string]string)
			if len(columns) == 0 {
				for _, col := range cols {
					row[col] = ""
				}
			} else {
				for _, col := range columns {
					row[col] = ""
				}
			}
			hasNonEmptyField := false
			for k, col := range data {
				if _, found := row[cols[k]]; found {
					row[cols[k]] = col.String
					fprintfWhileLenLessThanSpecifiedInt(sb, specLen, "%v:%v|", cols[k], col.String)
					if hasNonEmptyField || len(col.String) > 0 {
						hasNonEmptyField = true
					}
				}
			}
			ri++
			if !iter(row, hasNonEmptyField) {
				fmt.Fprintf(sb, "iter=0, further rows omitted.")
				break
			}
		}
	}
	if !logOnErrOnly {
		log.Detail(stage, sb.String())
	}
	return
}

func TxQuery(stage *log.Stage, tx *sql.Tx, s string, arguments ...interface{}) (resultInRows []map[string]string, err error) {
	stage.Enter("TxQuery")
	defer stage.Exit()
	log.Detail(stage, "[Tx]query[%v]args[%v]", s, arguments)

	rows, err := tx.Query(s, arguments...)
	if nil != err {
		log.Detail(stage, "[Tx]query[%v]args[%v]err:%v", s, arguments, err)
		return nil, err
	}

	sb := &strings.Builder{}
	fmt.Fprintf(sb, "[Tx]query[%v]args[%v]result:", s, arguments)
	specLen := sb.Len() + 1000
	if nil != rows {
		defer rows.Close()
		cols, _ := rows.Columns()
		buf := make([]interface{}, len(cols))
		data := make([]sql.NullString, len(cols))
		for i := range buf {
			buf[i] = &data[i]
		}

		ri := 0
		for rows.Next() {
			if err := rows.Scan(buf...); nil != err {
				log.Detail(stage, "[Tx]query[%v]args[%v]err:%v", s, arguments, err)
				return nil, err
			}
			fprintfWhileLenLessThanSpecifiedInt(sb, specLen, "[%v]", ri)
			row := make(map[string]string)

			for k, col := range data {
				row[cols[k]] = col.String
				fprintfWhileLenLessThanSpecifiedInt(sb, specLen, "%v:%v|", cols[k], col.String)
			}
			resultInRows = append(resultInRows, row)
		}
	}
	log.Detail(stage, sb.String())
	return
}

func TxExec(stage *log.Stage, tx *sql.Tx, s string, arguments ...interface{}) (result sql.Result, err error) {
	stage.Enter("TxExec")
	defer stage.Exit()
	log.Detail(stage, "[Tx]sql[%v]args[%v]", s, arguments)

	result, err = tx.Exec(s, arguments...)
	if nil != err {
		log.Detail(stage, "[Tx]sql[%v]args[%v]err:%v", s, arguments, err.Error())
		return nil, err

	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected() // error occurs only when driver doesn't support this function, just ignore
		log.Detail(stage, "[Tx]sql[%v]args[%v]Executed,RowsAffected[%v]", s, arguments, rowsAffected)
	}

	return result, nil
}

var fprintfWhileLenLessThanSpecifiedInt = func(sb *strings.Builder, len int, format string, a ...interface{}) {
	if sb.Len() >= len {
		return
	}
	fmt.Fprintf(sb, format, a...)
	if sb.Len() > len {
		t := sb.String()[:len]
		sb.Reset()
		sb.WriteString(t)
	}
}
