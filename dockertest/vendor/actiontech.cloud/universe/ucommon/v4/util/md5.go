package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	io_ "io"

	"actiontech.cloud/universe/ucommon/v4/io"
	"actiontech.cloud/universe/ucommon/v4/log"
)

type HashTimeoutError string

const HashingTimeoutPrefix = "[ERROR] Hashing timed out:"

func (e HashTimeoutError) Error() string {
	return fmt.Sprintf("%v %v", HashingTimeoutPrefix, string(e))
}

// IsHashTimeoutErr tells whether an error is a HashTimeoutError
func IsHashTimeoutErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), HashingTimeoutPrefix)
}

// MD5WithTimeout uses md5 algorithm to hash a file and returns a md5 checksum of that file
func MD5WithTimeout(stage *log.Stage, filePath string, timeoutSeconds int) (string, error) {
	stage.Enter("MD5WithTimeout")
	log.Detail(stage, "filepath: %s, timeoutSeconds:%v", filePath, timeoutSeconds)
	defer stage.Exit()

	if ok, _ := HasAutoQaDebugCondf("mock hashing time out error"); ok {
		return "", HashTimeoutError("mock md5 time out error")
	}

	file, err := os.Open(filePath)
	if nil != err {
		return "", fmt.Errorf("open %v error: %v", filePath, err)
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.CopyWithTimeout(hash, file, timeoutSeconds); err != nil {
		if _, ok := err.(io.ErrTimeout); ok {
			return "", HashTimeoutError("md5 time out error")
		} else {
			return "", fmt.Errorf("copy md5 hash error: %v", err)
		}
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// GetFileMd5 returns md5 checksum of a file
func GetFileMd5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if nil != err {
		return "", fmt.Errorf("open %v error: %v", filePath, err)
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io_.Copy(hash, file); err != nil {
		return "", fmt.Errorf("copy md5 hash error: %v", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
