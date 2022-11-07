package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"actiontech.cloud/universe/ucommon/v4/log"
)

const DebugTimeFormat = "2006-01-02 15:04:05"

var autoQaDebugCache string //performance

func ResetAutoQaDebugCache() {
	autoQaDebugCache = ""
}

func HasAutoQaDebugCondf(cond string, args ...interface{}) (ok bool, matches []string) {
	if "" == autoQaDebugCache {
		if IsFileExist("./auto_qa.debug") {
			autoQaDebugCache = "true"
		} else {
			autoQaDebugCache = "false"
		}
	}
	if "true" == autoQaDebugCache && IsFileExist("./auto_qa.debug") {
		if a, err := ioutil.ReadFile("./auto_qa.debug"); nil == err {
			return RegexMatch(string(a), fmt.Sprintf(cond, args...))
		}
	}
	return false, nil
}

func HasAutoQaDebugCondfWithAllMatches(cond string, args ...interface{}) (ok bool, matches [][]string) {
	if "" == autoQaDebugCache {
		if IsFileExist("./auto_qa.debug") {
			autoQaDebugCache = "true"
		} else {
			autoQaDebugCache = "false"
		}
	}
	if "true" == autoQaDebugCache && IsFileExist("./auto_qa.debug") {
		if a, err := ioutil.ReadFile("./auto_qa.debug"); nil == err {
			return AllRegexMatch(string(a), fmt.Sprintf(cond, args...))
		}
	}
	return false, nil
}

func HasAutoQaDebugCondfWithoutCache(cond string, args ...interface{}) (ok bool, matches []string) {
	if a, err := ioutil.ReadFile("./auto_qa.debug"); nil == err {
		return RegexMatch(string(a), fmt.Sprintf(cond, args...))
	}
	return false, nil
}

func DebugError(stage *log.Stage, cond string) error {
	ok, _ := HasAutoQaDebugCondf(cond)
	if ok {
		log.Key(stage, "[auto_qa.debug] "+cond)
		return fmt.Errorf("[auto_qa.debug] " + cond)
	}
	return nil
}

func DebugPause(cond string, args ...interface{}) {
	if ok, _ := HasAutoQaDebugCondf(cond, args...); !ok {
		return
	}
	log.Key(log.NewStage(), "DEBUG PAUSE %v", fmt.Sprintf(cond, args...))
	for {
		if ok, _ := HasAutoQaDebugCondf(cond, args...); ok {
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	log.Key(log.NewStage(), "DEBUG CONT %v", fmt.Sprintf(cond, args...))
}

func DebugPanic(a interface{}) {
	ok, _ := HasAutoQaDebugCondf(fmt.Sprintf("%v", a))
	if ok {
		log.Key(log.NewStage(), "[auto_qa.debug] %v", a)
		panic(a)
	}
}

// QADebugPause allows inserting a auto_qa.debug file in program's workdir to pause the program for
// given seconds. e.g., run `echo "pause for 10 seconds" > auto_qa.debug` to pause for 10 seconds.
func QADebugPause(stage *log.Stage) {
	if ok, match := HasAutoQaDebugCondf("pause for (\\d+) seconds"); ok {
		if len(match) > 1 {
			seconds, err := strconv.ParseInt(match[1], 10, 64)
			if err != nil {
				log.Detail(stage, "[ERROR] Failed to run QADebugPause:%s", err.Error())
				return
			}
			log.Detail(stage, "pause for %d seconds", seconds)
			time.Sleep(time.Second * time.Duration(seconds))
		}
	}
}

// QADebugTimeout enables using qa_debug file to set up timeoutValue for a given target
// e.g., set hashing timeout to be 1 seconds
func QADebugTimeout(stage *log.Stage, targetName string, timeout *int) error {
	text := fmt.Sprintf(`set %s timeout to be \[(\d+)\] seconds`, targetName)
	if ok, match := HasAutoQaDebugCondf(text); ok {
		timeoutStr := match[1]
		if timeoutVal, err := strconv.Atoi(timeoutStr); err != nil {
			return err
		} else {
			log.Key(stage, "[auto_qa.debug] set %s timeout to be [%v] seconds", targetName, timeout)
			timeout = &timeoutVal
		}
	}
	return nil
}
