package wrapper

import "time"

const wrapperLogTimestamp = "2006-01-02 15:04:05.000"

func GetCurrentTimestamp() string {
	return time.Now().Format(wrapperLogTimestamp)
}
