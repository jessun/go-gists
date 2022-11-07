package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pkg/errors"

	ucommonLog "actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/mysql"
	"actiontech.cloud/universe/ucommon/v4/util"
)

var unitTestDB *mysql.DB
var unitTestRawDB *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		err = errors.Wrap(err, "unit_test: failed to create docker pool")
		log.Fatalln(err.Error())
	}

	log.Println("unit_test: create docker pool successfully")

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "oceanbase/oceanbase-ce",
		Name:         "test_mysql",
		Tag:          "3.1.4",
		Env:          []string{"MINI_MODE=1"},
		ExposedPorts: []string{"2881"},
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
	})

	if err != nil {
		err = errors.Wrap(err, "unit_test: failed to run mysql docker")
		log.Fatalln(err.Error())
	}

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	log.Println("unit_test: run oceanbase-ce image successfully")
	port := resource.GetPort("2881/tcp")
	addr := resource.GetHostPort("2881/tcp")
	log.Printf("unit_test: oceanbase-ce port: %s, addr: %s", port, addr)

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	// we will receive an error "packets.go:37: unexpected EOF " if we don't retry
	// "[mysql] 2022/10/06 21:40:53 packets.go:37: unexpected EOF"
	ok := <-util.WaitUntil(func() bool {
		var err error
		unitTestRawDB, err = sql.Open("mysql", fmt.Sprintf("root@(localhost:%s)/", resource.GetPort("2881/tcp")))
		if err != nil {
			err = errors.Wrap(err, "unit_test: failed to open mysql connection")
			log.Println(err.Error())
			return false
		}

		if err := unitTestRawDB.Ping(); err != nil {
			err = errors.Wrap(err, "unit_test: failed to ping mysql")
			log.Println(err.Error())
			return false
		}
		return true
	}, 10*60)
	if !ok {
		log.Fatalln("unit_test: failed to wait until mysql is ready")
	}

	unitTestDB = mysql.NewDbWithoutCache(unitTestRawDB, 10)

	code := m.Run()
	defer os.Exit(code)
}

func TestSomething(t *testing.T) {
	stage := ucommonLog.NewStage().Enter("TestSomething")
	defer stage.Exit()
	ok := <-util.WaitUntil(func() bool {
		res, err := mysql.SqlQuery(stage, unitTestDB, "SHOW DATABASES")
		if err != nil {
			err = errors.Wrap(err, "mysql query failed")
			log.Println(err.Error())
			return false
		}
		log.Printf("databases: [%#v]", res)

		return false
	}, 10*60)

	if !ok {
		log.Print("TestSomething failed")
	}

}
