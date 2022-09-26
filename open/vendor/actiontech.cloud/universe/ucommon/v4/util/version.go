package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Version represents any software's version number that consists of three parts.
type Version struct {
	raw      string
	segments [3]uint64
}

func NewVersionWithNumbers(major, minor, revision uint64) Version {
	return Version{
		raw:      fmt.Sprintf("%d.%d.%d", major, minor, revision),
		segments: [3]uint64{major, minor, revision},
	}
}

func NewVersion(s string) (Version, error) {
	v := Version{
		raw: s,
	}
	segments := strings.Split(s, ".")
	if len(segments) != 3 {
		return Version{}, errors.New("[ERROR] Version should be three parts")
	}

	for i := 0; i < 3; i++ {
		num, err := strconv.ParseUint(segments[i], 10, 64)
		if err != nil {
			return Version{}, fmt.Errorf("[ERROR] Failed to parse version(%s):%s", s, err.Error())
		}
		v.segments[i] = num
	}

	return v, nil
}

func (m Version) NoLessThan(ver2 Version) bool {
	segments1 := m.segments
	segments2 := ver2.segments
	for i := 0; i < 3; i++ {
		if segments1[i] != segments2[i] {
			return segments1[i] > segments2[i]
		}
	}
	return true
}

func (m Version) LessThan(ver2 Version) bool {
	return !m.NoLessThan(ver2)
}

func (m Version) MainVersion() string {
	return fmt.Sprintf("%v.%v", m.segments[0], m.segments[1])
}

// IsInTheRange 判断MySQL版本号是否在某个区间内(包含传入的区间)
func (m Version) IsInTheRange(min, max Version) bool {
	var checkResult = true
	for i := 0; i < 3; i++ {
		if !(min.segments[i] <= m.segments[i] && m.segments[i] <= max.segments[i]) {
			checkResult = false
			break
		}
	}
	return checkResult
}

func (m Version) HasNewerRevision(newVersion Version) bool {
	if len(m.segments) != len(newVersion.segments) {
		return false
	}
	if m.segments[0] == newVersion.segments[0] && m.segments[1] == newVersion.segments[1] {
		return m.segments[2] < newVersion.segments[2]
	}
	return false
}

func (m Version) String() string {
	return m.raw
}
