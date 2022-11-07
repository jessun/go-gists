// reference: https://github.com/gyuho/linux-inspect/tree/master/df
package os

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

// Row is 'df' command output row in Linux.
type Row struct {
	// FileSystem is file system ('source').
	FileSystem string `column:"file_system"`
	// Device is device name.
	Device string `column:"device"`
	// MountedOn is 'mounted on' ('target').
	MountedOn string `column:"mounted_on"`
	// FileSystemType is file system type ('fstype').
	FileSystemType string `column:"file_system_type"`
	// File is file name if specified on the command line ('file').
	File string `column:"file"`
	// Inodes is total number of inodes ('itotal').
	Inodes int64 `column:"inodes"`
	// Ifree is number of available inodes ('iavail').
	Ifree int64 `column:"ifree"`
	// Iused is number of used inodes ('iused').
	Iused int64 `column:"iused"`
	// IusedPercent is percentage of iused divided by itotal ('ipcent').
	IusedPercent string `column:"iused_percent"`
	// TotalBlocks is total number of 1K-blocks ('size').
	TotalBlocks            int64  `column:"total_blocks"`
	TotalBlocksBytesN      int64  `column:"total_blocks_bytes_n"`
	TotalBlocksParsedBytes string `column:"total_blocks_parsed_bytes"`
	// AvailableBlocks is number of available 1K-blocks ('avail').
	AvailableBlocks            int64  `column:"available_blocks"`
	AvailableBlocksBytesN      int64  `column:"available_blocks_bytes_n"`
	AvailableBlocksParsedBytes string `column:"available_blocks_parsed_bytes"`
	// UsedBlocks is number of used 1K-blocks ('used').
	UsedBlocks            int64  `column:"used_blocks"`
	UsedBlocksBytesN      int64  `column:"used_blocks_bytes_n"`
	UsedBlocksParsedBytes string `column:"used_blocks_parsed_bytes"`
	// UsedBlocksPercent is percentage of used-blocks divided by total-blocks ('pcent').
	UsedBlocksPercent string `column:"used_blocks_percent"`
}

// dfPath is the default 'df' command path.
const dfPath = "/bin/df"

// Get returns entries in 'df' command.
// Pass '' target to list all information.
func GetDfInfo(dfPath string, target string) ([]Row, error) {
	o, err := Read(dfPath, target)
	if err != nil {
		return nil, err
	}
	return ParseDf(o)
}

// GetDefault returns entries in 'df' command.
// Pass '' target to list all information.
func GetDfDefault(target string) ([]Row, error) {
	o, err := Read(dfPath, target)
	if err != nil {
		return nil, err
	}
	return ParseDf(o)
}

// Read reads Linux 'df' command output.
// Pass '' target to list all information.
func Read(dfPath string, target string) (string, error) {
	buf := new(bytes.Buffer)
	err := readDfInfo(dfPath, target, buf)
	o := strings.TrimSpace(buf.String())
	return o, err
}

func readDfInfo(dfPath string, target string, w io.Writer) error {
	if !IsFileExist(dfPath) {
		return fmt.Errorf("%q does not exist", dfPath)
	}
	// dfFlags is 'df --all --sync --block-size=1024 --output=source,target,fstype,file,itotal,iavail,iused,ipcent,size,avail,used,pcent'.
	var dfFlags = []string{"--all", "--sync", "--block-size=1024", "--output=source,target,fstype,file,itotal,iavail,iused,ipcent,size,avail,used,pcent"}

	if target != "" {
		dfFlags = append(dfFlags, strings.TrimSpace(target))
	}
	cmd := exec.Command(dfPath, dfFlags...)
	cmd.Stdout = w
	cmd.Stderr = w
	return cmd.Run()
}

// Headers is the headers in 'df' output.
var Headers = []string{
	"Filesystem",

	// Mounted on
	"Mounted",
	"on",

	"Type",
	"File",
	"Inodes",
	"IFree",
	"IUsed",
	"IUse%",
	"1K-blocks",
	"Avail",
	"Used",
	"Use%",
}

type commandOutpudrowIdx int

const (
	command_output_row_idx_file_system commandOutpudrowIdx = iota
	command_output_row_idx_mounted_on
	command_output_row_idx_file_system_type
	command_output_row_idx_file
	command_output_row_idx_inodes
	command_output_row_idx_ifree
	command_output_row_idx_iused
	command_output_row_idx_iused_percent
	command_output_row_idx_total_blocks
	command_output_row_idx_available_blocks
	command_output_row_idx_used_blocks
	command_output_row_idx_used_blocks_percentage
)

// Parse parses 'df' command output and returns the rows.
func ParseDf(s string) ([]Row, error) {
	lines := strings.Split(s, "\n")
	rows := make([][]string, 0, len(lines))
	headerFound := false
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		ds := strings.Fields(strings.TrimSpace(line))
		if ds[0] == "Filesystem" { // header line
			if !reflect.DeepEqual(ds, Headers) {
				return nil, fmt.Errorf("unexpected 'df' command header order (%v, expected %v, output: %q)", ds, Headers, s)
			}
			headerFound = true
			continue
		}

		if !headerFound {
			continue
		}

		row := strings.Fields(strings.TrimSpace(line))
		if len(row) != len(Headers)-1 {
			return nil, fmt.Errorf("unexpected row column number %v (expected %v)", row, Headers)
		}
		rows = append(rows, row)
	}

	type result struct {
		row Row
		err error
	}
	rc := make(chan result, len(rows))
	for _, row := range rows {
		go func(row []string) {
			tr, err := parseDfRow(row)
			rc <- result{row: tr, err: err}
		}(row)
	}

	tcRows := make([]Row, 0, len(rows))
	for len(tcRows) != len(rows) {
		select {
		case rs := <-rc:
			if rs.err != nil {
				return nil, rs.err
			}
			tcRows = append(tcRows, rs.row)
		}
	}
	rm := make(map[string]Row)
	for _, row := range tcRows {
		rm[row.MountedOn] = row
	}
	rrs := make([]Row, 0, len(rm))
	for _, row := range rm {
		rrs = append(rrs, row)
	}
	return rrs, nil
}

func parseDfRow(row []string) (Row, error) {
	drow := Row{
		FileSystem:        strings.TrimSpace(row[command_output_row_idx_file_system]),
		MountedOn:         strings.TrimSpace(row[command_output_row_idx_mounted_on]),
		FileSystemType:    strings.TrimSpace(row[command_output_row_idx_file_system_type]),
		File:              strings.TrimSpace(row[command_output_row_idx_file]),
		IusedPercent:      strings.TrimSpace(strings.Replace(row[command_output_row_idx_iused_percent], "%", " %", -1)),
		UsedBlocksPercent: strings.TrimSpace(strings.Replace(row[command_output_row_idx_used_blocks_percentage], "%", " %", -1)),
	}
	drow.Device = filepath.Base(drow.FileSystem)

	ptxt := strings.TrimSpace(row[command_output_row_idx_inodes])
	if ptxt == "-" {
		ptxt = "0"
	}
	iv, err := strconv.ParseInt(ptxt, 10, 64)
	if err != nil {
		return Row{}, fmt.Errorf("parse error %v (row %v)", err, row)
	}
	drow.Inodes = iv

	ptxt = strings.TrimSpace(row[command_output_row_idx_ifree])
	if ptxt == "-" {
		ptxt = "0"
	}
	iv, err = strconv.ParseInt(ptxt, 10, 64)
	if err != nil {
		return Row{}, fmt.Errorf("parse error %v (row %v)", err, row)
	}
	drow.Ifree = iv

	ptxt = strings.TrimSpace(row[command_output_row_idx_iused])
	if ptxt == "-" {
		ptxt = "0"
	}
	iv, err = strconv.ParseInt(ptxt, 10, 64)
	if err != nil {
		return Row{}, fmt.Errorf("parse error %v (row %v)", err, row)
	}
	drow.Iused = iv

	ptxt = strings.TrimSpace(row[command_output_row_idx_total_blocks])
	if ptxt == "-" {
		ptxt = "0"
	}
	iv, err = strconv.ParseInt(ptxt, 10, 64)
	if err != nil {
		return Row{}, fmt.Errorf("parse error %v (row %v)", err, row)
	}
	drow.TotalBlocks = iv
	drow.TotalBlocksBytesN = iv * 1024
	drow.TotalBlocksParsedBytes = humanize.Bytes(uint64(drow.TotalBlocksBytesN))

	ptxt = strings.TrimSpace(row[command_output_row_idx_available_blocks])
	if ptxt == "-" {
		ptxt = "0"
	}
	iv, err = strconv.ParseInt(ptxt, 10, 64)
	if err != nil {
		return Row{}, fmt.Errorf("parse error %v (row %v)", err, row)
	}
	drow.AvailableBlocks = iv
	drow.AvailableBlocksBytesN = iv * 1024
	drow.AvailableBlocksParsedBytes = humanize.Bytes(uint64(drow.AvailableBlocksBytesN))

	ptxt = strings.TrimSpace(row[command_output_row_idx_used_blocks])
	if ptxt == "-" {
		ptxt = "0"
	}
	iv, err = strconv.ParseInt(ptxt, 10, 64)
	if err != nil {
		return Row{}, fmt.Errorf("parse error %v (row %v)", err, row)
	}
	drow.UsedBlocks = iv
	drow.UsedBlocksBytesN = iv * 1024
	drow.UsedBlocksParsedBytes = humanize.Bytes(uint64(drow.UsedBlocksBytesN))

	return drow, nil
}
