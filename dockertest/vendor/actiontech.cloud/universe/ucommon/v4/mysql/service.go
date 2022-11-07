package mysql

import (
	"strconv"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/os"
	"actiontech.cloud/universe/ucommon/v4/scsi"
)

func GetSystemdVersion(stage *log.Stage) (int, error) {
	ret, err := os.Cmdf2(stage, "SUDO systemctl --version | head -1 | awk '{print $2;}'")
	if err != nil {
		log.Detail(stage, "get systemd version error: [%v]", err)
		return 0, err
	}
	systemdVersion, err := strconv.Atoi(ret)
	if err != nil {
		log.Detail(stage, "convert systemd version error: [%v]", err)
		return 0, err
	}
	return systemdVersion, nil
}

func IsMysqlServiceAlive(stage *log.Stage, mysqlPort int, mysqlRunUser string) bool {
	status := false
	if os.IsSystemd() {
		_, err := os.Cmdf2(stage, "SUDO systemctl status mysqld_%v.service", mysqlPort)
		status = nil == err
	} else {
		_, err := os.Cmdf2(stage, `[ -e /etc/init.d/mysqld_%v ] && SU %v -c "/etc/init.d/mysqld_%v status"`, mysqlPort, mysqlRunUser, mysqlPort)
		status = nil == err
	}
	return status
}

func StartMysqlService(stage *log.Stage, mysqlPort int, mysqlRunUser, scsiDev, scsiPr, scsiMountpoint string, scsiPrLevel int) error {
	stage.Enter("start_mysql_service")
	defer stage.Exit()

	if IsMysqlServiceAlive(stage, mysqlPort, mysqlRunUser) {
		return nil
	}

	if "" != scsiDev {
		if err := <-scsi.ReserveOrPreempt(stage, scsiDev, scsiPr, scsiPrLevel, false); nil != err {
			return err
		}
		os.Umount(stage, scsiMountpoint)
		if err := os.Mount(stage, scsiDev, scsiMountpoint, false, mysqlRunUser); nil != err {
			return err
		}
	}

	if os.IsSystemd() {
		if _, err := os.Cmdf2(stage, "SUDO systemctl start mysqld_%v.service", mysqlPort); nil != err {
			return err
		}
	} else {
		_, err := os.Cmdf2(stage, `[ -e /etc/init.d/mysqld_%v ] && SU %v -c "/etc/init.d/mysqld_%v start"`,
			mysqlPort, mysqlRunUser, mysqlPort)
		if nil != err {
			return err
		}
	}

	return nil
}

func StopMysqlService(stage *log.Stage, mysqlPort int, mysqlRunUser string) error {
	stage.Enter("stop_mysql_service")
	defer stage.Exit()

	if !IsMysqlServiceAlive(stage, mysqlPort, mysqlRunUser) {
		return nil
	}

	if os.IsSystemd() {
		if _, err := os.Cmdf2(stage, "SUDO systemctl stop mysqld_%v.service", mysqlPort); nil != err {
			return err
		}

	} else {
		_, err := os.Cmdf2(stage, `[ -e /etc/init.d/mysqld_%v ] && SU %v -c "/etc/init.d/mysqld_%v stop"`,
			mysqlPort, mysqlRunUser, mysqlPort)
		if nil != err {
			return err
		}
	}

	return nil
}
