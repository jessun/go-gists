package util

import (
	"fmt"
	"strings"
	"time"
)

const formatErrorTemplate = "format parsing failed,err: %v |expected format: %v"
const FormatMistaken = "format mistaken"

// 这是个大坑, time.Parse()的时区信息是 UTC, time.Now()取到的是CST, 这会导致两个时间差了8个小时,从而导致计算时间差时出错
var TimeLocationCST = time.FixedZone("CST", 8*3600)

type FormatError struct {
	err     interface{}
	example string
}

func (f *FormatError) Error() string {
	return fmt.Sprintf(formatErrorTemplate, f.err, f.example)
}

func NewFormatError(err interface{}, example string) error {
	return &FormatError{
		err:     err,
		example: example,
	}
}

const TimeParseStr = "2006-01-02-15" // processlist 需求是精确到小时

type ProcesslistLogNameFormat struct {
	MysqlInstanceID string    // mysql实例ID
	Time            time.Time // 日志生成时间
}

func (p *ProcesslistLogNameFormat) Format() (string, error) {
	tm := p.Time.Format(TimeParseStr)
	return fmt.Sprintf("%v_%v.csv", tm, p.MysqlInstanceID), nil
}

func (p *ProcesslistLogNameFormat) Parse(str string) error {
	fns := strings.Split(str, "_")
	if len(fns) != 2 {
		return NewFormatError(FormatMistaken, p.Example())
	}
	su := strings.Split(fns[1], ".")
	if len(su) != 2 || su[1] != "csv" {
		return NewFormatError(FormatMistaken, p.Example())
	}
	tm, err := time.ParseInLocation(TimeParseStr, fns[0], TimeLocationCST)
	if err != nil {
		return NewFormatError(err, p.Example())
	}
	p.MysqlInstanceID = su[0]
	p.Time = tm
	return nil
}

func (p *ProcesslistLogNameFormat) Example() string {
	return "2006-01-02-15_mysql-id.csv"
}

type ProcesslistLogArchiveNameFormat struct {
	MysqlInstanceID string    // mysql实例ID
	Time            time.Time // 日志生成时间
}

func (p *ProcesslistLogArchiveNameFormat) Format() (string, error) {
	tm := p.Time.Format(TimeParseStr)
	return fmt.Sprintf("%v_%v.7z", tm, p.MysqlInstanceID), nil
}

func (p *ProcesslistLogArchiveNameFormat) Parse(str string) error {
	fns := strings.Split(str, "_")
	if len(fns) != 2 {
		return NewFormatError(FormatMistaken, p.Example())
	}
	su := strings.Split(fns[1], ".")
	if len(su) != 2 || su[1] != "7z" {
		return NewFormatError(FormatMistaken, p.Example())
	}
	tm, err := time.ParseInLocation(TimeParseStr, fns[0], TimeLocationCST)
	if err != nil {
		return NewFormatError(err, p.Example())
	}
	p.MysqlInstanceID = su[0]
	p.Time = tm
	return nil
}

func (p *ProcesslistLogArchiveNameFormat) Example() string {
	return "2006-01-02-15_mysql-id.7z"
}
