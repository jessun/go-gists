package mysql

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/os"
)

const MyCNFPrefix = "my.cnf"
const Dot = "."

// ParseVersionFromMysqldCmd parse single mysql version from "mysqld --version" commond output.
func ParseVersionFromMysqldCmd(cmdVersion string) (string, error) {
	segs := regexp.MustCompile("Ver +([^ \\-]+)").FindStringSubmatch(cmdVersion)
	if nil == segs {
		return "", fmt.Errorf("can't match mysqld version")
	} else {
		return segs[1], nil
	}
}

// GetMysqlOneDigitVersion returns 5 for mysqlVersion 5.7.25
func GetMysqlOneDigitVersion(mysqlVersion string) string {
	versionNumbers := strings.Split(mysqlVersion, Dot)
	if len(versionNumbers) == 0 {
		return ""
	}
	return versionNumbers[0]
}

// GetMysqlTwoDigitsVersion returns 5.7 for mysqlVersion 5.7.25
func GetMysqlTwoDigitsVersion(mysqlVersion string) string {
	versionNumbers := strings.Split(mysqlVersion, Dot)
	if len(versionNumbers) < 2 {
		return ""
	}
	return strings.Join(versionNumbers[:2], Dot)
}

// GetMyCNFName returns MySQL's my.cnf filename
// the second parameter is myCnfSuffix, "udb","group_replication", etc. Leave it empty if there is none.
func GetMyCNFName(mysqlVersion, myCnfSuffix string) string {
	myCnfStrings := []string{MyCNFPrefix}
	twoDigitsVersion := GetMysqlTwoDigitsVersion(mysqlVersion)
	if twoDigitsVersion != "" {
		myCnfStrings = append(myCnfStrings, twoDigitsVersion)
	}
	if myCnfSuffix != "" {
		myCnfStrings = append(myCnfStrings, myCnfSuffix)
	}
	return strings.Join(myCnfStrings, Dot)
}

// GetMysqlVersion returns local MySQL daemon's version
func GetMysqlVersion(stage *log.Stage, baseDir string) (string, error) {
	mysqldPath := filepath.Join(baseDir, "bin", "mysqld")
	out, err := os.Cmdf2(stage, "%v --version", mysqldPath)
	if nil != err {
		return "", fmt.Errorf("%v --version, error: %v", mysqldPath, err)
	}
	v, err := ParseVersionFromMysqldCmd(out)
	if nil != err {
		return "", fmt.Errorf("%v --version, error: %v", mysqldPath, err)
	}
	return v, nil
}
