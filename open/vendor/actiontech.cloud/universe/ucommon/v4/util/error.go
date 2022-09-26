package util

import "strings"

type ErrList []error

func (el ErrList) Error() string {
	errMsg := make([]string, len(el))

	for i := range el {
		errMsg[i] = el[i].Error()
	}

	return strings.Join(errMsg, ", ")
}
