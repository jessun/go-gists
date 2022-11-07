package util

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"actiontech.cloud/universe/ucommon/v4/log"
	xxhash "github.com/cespare/xxhash/v2"
)

// XXHashWithTimeout wraps XXHash with timeout feature, if after timeout seconds XXHash fails
// to compute a hash value, it returns a timeout error.
func XXHashWithTimeout(stage *log.Stage, filePath string, timeoutSeconds int) (string, error) {
	stage.Enter("HashWithTimeout")
	defer stage.Exit()

	timer := time.NewTimer(time.Second * time.Duration(timeoutSeconds))
	jobDone := make(chan interface{})
	defer close(jobDone)
	go func(stage *log.Stage) {
		sum, err := XXHash(stage, filePath)
		if err != nil {
			jobDone <- err
		}
		QADebugPause(stage) // to fake doing hash calculation for a given seconds
		jobDone <- sum
	}(stage.Go())

	select {
	case <-timer.C:
		log.Key(stage, "[ERROR] Hash backup file (%s) timeout after:%v seconds", filePath, timeoutSeconds)
		return "", HashTimeoutError(fmt.Sprintf("file:%s, timeout seconds:%v", filePath, timeoutSeconds))
	case result := <-jobDone:
		if value, isErr := result.(error); isErr {
			log.Key(stage, "[ERROR] Hash backup file (%s) failed: %s", filePath, value)
			return "", value
		}
		return result.(string), nil
	}
}

// XXHash is an extremely fast non-cryptographic hash algorithm
func XXHash(stage *log.Stage, filepath string) (string, error) {
	stage.Enter("XXHash")
	defer stage.Exit()

	f, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("[ERROR] XXHash: Failed to open file(%s):%s", filepath, err.Error())
	}
	defer f.Close()

	h := xxhash.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("[ERROR] XXHash: Failed to io.Copy file(%s):%s", filepath, err.Error())
	}
	checkSum := strconv.FormatUint(h.Sum64(), 10)
	log.Detail(stage, "Successfully calculated checksum(%s) of file(%s)", checkSum, filepath)
	return checkSum, nil
}
