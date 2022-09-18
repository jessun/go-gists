package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `recovery_window=7d&auto_delete_obsolete_backup=true&log_archive_piece_switch_interval=1d`
	ss := strings.Split(s, "&")

	for i := range ss {
		opt := ss[i]
		optKeyVal := strings.Split(opt, "=")
		fmt.Printf("KEY: %v, VAL: %v", optKeyVal[0], optKeyVal[1])
	}
}
