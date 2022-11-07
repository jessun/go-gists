package util

import (
	"bufio"
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime/pprof"
	"strconv"
	"strings"
	builtin_syscall "syscall"
	"time"

	"actiontech.cloud/universe/ucommon/v4/conf"
	"actiontech.cloud/universe/ucommon/v4/log"
)

// OBSOLETE, StringNvl instead
func Nvl(args ...string) string {
	for _, arg := range args {
		if "" != arg {
			return arg
		}
	}
	return ""
}

func ErrorNvl(errs ...error) error {
	for _, err := range errs {
		if nil != err {
			return err
		}
	}
	return nil
}

func Ping(db *sql.DB) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered in db.Ping():%v", r)
		}
	}()
	err = db.Ping()
	return err
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return nil == err
}

func IsFileNotExist(path string) bool {
	_, err := os.Stat(path)
	return nil != err && os.IsNotExist(err)
}

func IsDir(path string) (bool, error) {
	if s, err := os.Stat(path); nil != err {
		return false, err
	} else {
		return s.IsDir(), nil
	}
}

func IsEmptyDir(path string) bool {
	infos, err := ioutil.ReadDir(path)
	if nil != err {
		return true
	}
	return 0 == len(infos)
}

func IsEmptyDirExceptHa(path string) bool {
	infos, err := ioutil.ReadDir(path)
	if nil != err {
		return true
	}
	for _, info := range infos {
		if !strings.HasPrefix(info.Name(), "HA_") {
			return false
		}
	}
	return true
}

func RegexMatch(input string, regexStr string) (ok bool, matches []string) {
	regex := regexp.MustCompile(regexStr)
	if !regex.MatchString(input) {
		return false, nil
	} else {
		return true, regex.FindStringSubmatch(input)
	}
}

func AllRegexMatch(input string, regexStr string) (ok bool, matches [][]string) {
	regex := regexp.MustCompile(regexStr)
	if !regex.MatchString(input) {
		return false, nil
	} else {
		return true, regex.FindAllStringSubmatch(input, -1)
	}
}

func Atob(a string) byte {
	i, _ := strconv.Atoi(a)
	return byte(i)
}

func Uniq(arr []string) []string {
	ret := make([]string, 0)
	for _, ele := range arr {
		if Contains(ret, ele) {
			goto DUP
		}
		ret = append(ret, ele)
	DUP:
	}
	return ret
}

func SplitAndTrimSpace(s string, spliter string) []string {
	return SplitAndTrimSpaceN(s, spliter, -1)
}

func SplitAndTrimSpaceN(s string, spliter string, n int) []string {
	ret := make([]string, 0)
	for _, ele := range strings.SplitN(s, spliter, n) {
		if "" != ele {
			ret = append(ret, strings.TrimSpace(ele))
		}
	}
	return ret
}

func GoAndWaitOnePass(ips []string, fn func(string) bool) bool {
	retChan := make(chan bool, len(ips))
	for _, ip := range ips {
		go func(ip string) {
			retChan <- fn(ip)
		}(ip)
	}
	for _ = range ips {
		if <-retChan {
			return true
		}
	}
	return false
}

func All(ips []string, fn func(string) bool) bool {
	retChan := make(chan bool, len(ips))
	for _, ip := range ips {
		go func(ip string) {
			retChan <- fn(ip)
		}(ip)
	}
	for _ = range ips {
		if !<-retChan {
			return false
		}
	}
	return true
}

func Any(ips []string, fn func(string) bool) bool {
	return GoAndWaitOnePass(ips, fn)
}

func BytesToUint64(buf []byte) uint64 { //NO_DUPL_CHECK
	var a uint64
	var i uint
	for _, b := range buf {
		a += uint64(b) << i
		i += 8
	}
	return a
}

func BytesToUint(buf []byte) uint {
	var a uint
	var i uint
	for _, b := range buf {
		a += uint(b) << i
		i += 8
	}
	return a
}

func UintToBytes(num uint, buf []byte) []byte { //NO_DUPL_CHECK
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(num & 0xff)
		num = num >> 8
	}
	return buf
}

func Uint64ToBytes(num uint64, buf []byte) []byte {
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(num & 0xff)
		num = num >> 8
	}
	return buf
}

func GetBinlogSize(binlogDir string, firstBinlogName string, firstBinlogStartPosStr string /*int64*/) (int64, error) {
	if "" == firstBinlogName {
		return 0, nil
	}
	var firstBinlogStartPos int64 = 0
	if "" != firstBinlogStartPosStr {
		if a, err := strconv.ParseInt(firstBinlogStartPosStr, 10, 64); nil != err {
			return 0, err
		} else {
			firstBinlogStartPos = a
		}
	}

	var totalSize int64 = 0
	err := ForEachBinlogInDir(binlogDir, firstBinlogName, func(info os.FileInfo) error {
		if info.Name() == firstBinlogName {
			totalSize += info.Size() - firstBinlogStartPos
		} else {
			totalSize += info.Size()
		}
		return nil
	})

	return totalSize, err
}

var BinlogFileNamePattern = regexp.MustCompile(`^[^\.]+[\.]\d{6}$`)

func GetBinlogPrefix(binlog string) string {
	if !BinlogFileNamePattern.MatchString(binlog) {
		return binlog
	}
	return binlog[:len(binlog)-7]
}

func ForEachBinlogInDir(binlogDir string, firstBinlogName string, fn func(os.FileInfo) error) error {
	binlogFileNameBase := GetBinlogPrefix(firstBinlogName)
	fileInfos, err := ioutil.ReadDir(binlogDir)
	if nil != err {
		return err
	}
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if strings.HasPrefix(fileName, binlogFileNameBase) && fileName >= firstBinlogName && BinlogFileNamePattern.MatchString(fileName) {
			err := fn(fileInfo)
			if nil != err {
				return err
			}
		}
	}
	return nil
}

func GetFilesSize(files []string) (int64, error) {
	var totalSize int64
	for _, file := range files {
		stat, err := os.Stat(file)
		if nil != err {
			return 0, err
		}
		totalSize += stat.Size()
	}
	return totalSize, nil
}

func GetDirFilesSize(dirpath string) (uint64, error) {
	var binlogSize uint64
	fileInfos, err := ioutil.ReadDir(dirpath)
	if nil != err {
		return uint64(0), err
	}
	for _, fi := range fileInfos {
		if fi.IsDir() {
			continue
		}
		binlogSize += uint64(fi.Size())
	}
	return binlogSize, nil
}

func Prop(s, p string) (result string) {
	if reg := regexp.MustCompile(fmt.Sprintf(`%s\s*=\s*(\S*)\s*`, p)).FindStringSubmatch(s); len(reg) > 1 {
		result = reg[len(reg)-1]
	}
	return result
}

func DumpLoop(rootDir string) {
	c := make(chan os.Signal, 10)
	signal.Notify(c, builtin_syscall.Signal(0x15), builtin_syscall.Signal(0xc)) //0x15=SIGTTIN 0xc=SIGUSR2

	for {
		sig := <-c
		switch sig {
		case builtin_syscall.Signal(0x15):
			go func() {
				pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
				if f, err := os.OpenFile(filepath.Join(rootDir, "dump"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0640); nil != err { //NO_DUPL_CHECK
					fmt.Fprintf(os.Stderr, "write dump error(%v)", err)
				} else {
					pprof.Lookup("goroutine").WriteTo(f, 1)
					f.Close()
				}
				{
					f, err := os.OpenFile(filepath.Join(rootDir, "heap_dump"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0640)
					if nil != err {
						fmt.Fprintf(os.Stderr, "write heap_dump error(%v)", err)
					}
					pprof.WriteHeapProfile(f)
					f.Close()
				}
				return
			}()
		case builtin_syscall.Signal(0xc):
			go func() {
				f, err := os.Create(filepath.Join(rootDir, "cpu_profile"))
				if err != nil {
					fmt.Fprintf(os.Stderr, "write cpu_profile error(%v)", err)
				}
				defer f.Close()
				pprof.StartCPUProfile(f)
				time.Sleep(5 * time.Minute)
				pprof.StopCPUProfile()
			}()
		default:
		}

	}
}

func StopChanLoop() chan bool {
	c := make(chan os.Signal, 10)
	ret := make(chan bool)
	signal.Notify(c, builtin_syscall.SIGUSR1)
	go func() {
		for {
			select {
			case <-c:
				ret <- true
			}
		}
	}()
	return ret
}

func TimeoutChan(seconds int) chan bool {
	ret := make(chan bool, 1)
	go func() {
		select {
		case <-time.After(time.Duration(seconds) * time.Second):
			ret <- true
		}
	}()
	return ret
}

func GenFileNameByTime() string {
	now := time.Now().Format("2006-01-02_15-04-05.000000000") // means yyyy-MM-dd_HH-mm-ss-nanoSeconds
	return strings.Replace(now, ".", "-", -1)                 //nanoSeconds is required to be lead by ".". replace "." to "-"
}

var LengthOfFileNameByTime = len(GenFileNameByTime())
var PatternOfFileNameByTime = regexp.MustCompile("\\d\\d\\d\\d\\-\\d\\d\\-\\d\\d\\_\\d\\d\\-\\d\\d\\-\\d\\d\\-\\d\\d\\d\\d\\d\\d\\d\\d\\d$")

func ParseTimeByFileName(timestamp string) time.Time {
	timestamp = timestamp[0:19] + "." + timestamp[20:]
	t, _ := time.ParseInLocation("2006-01-02_15-04-05.000000000", timestamp, time.Local)
	return t
}

func AddConfigSectionToConfig(src *conf.ConfigFile, section string, target *conf.ConfigFile) error {
	if !src.HasSection(section) {
		return fmt.Errorf("no section (%v) found in src conf", section)
	}
	if target.HasSection(section) {
		return fmt.Errorf("section (%v) already exist in target conf", section)
	}
	options, err := src.GetOptions(section)
	if nil != err {
		return err
	}
	target.AddSection(section)
	for _, option := range options {
		val, _ := src.GetString(section, option)
		target.AddOption(section, option, val)
	}
	return nil
}

func UpdateConfigSectionToConfig(src *conf.ConfigFile, section string, target *conf.ConfigFile) (delta [][]string, err error) {
	if !src.HasSection(section) {
		return nil, fmt.Errorf("no section (%v) found in src conf", section)
	}
	if !target.HasSection(section) {
		return nil, fmt.Errorf("no section (%v) found in target conf", section)
	}
	options, err := src.GetOptions(section)
	if nil != err {
		return nil, err
	}
	delta = make([][]string, 0)
	for _, option := range options {
		val, _ := src.GetString(section, option)
		oldVal, _ := target.GetString(section, option)
		if "" == val {
			target.RemoveOption(section, option)
		} else {
			target.AddOption(section, option, val)
		}
		if oldVal != val {
			delta = append(delta, []string{section, option, oldVal, val})
		}
	}
	return delta, nil
}

func GetFirstNonDefaultSectionInConf(c *conf.ConfigFile) string {
	for _, section := range c.GetSections() {
		if "default" != section {
			return section
		}
	}
	return ""
}

func Contains(arr []string, ele string) bool {
	for _, a := range arr {
		if ele == a {
			return true
		}
	}
	return false
}

func RemoveAll(arr []string, ele string) (result []string) {
	for _, v := range arr {
		if !strings.Contains(v, ele) {
			result = append(result, v)
		}
	}
	return
}

// BatchOverlap is used when getting slices' intersections within a for loop
type BatchOverlap struct {
	strings  [][]string
	hasEmpty bool
}

func (b *BatchOverlap) Add(s []string) *BatchOverlap {
	if b.hasEmpty {
		return b
	}
	if len(s) == 0 {
		b.hasEmpty = true
	}
	b.strings = append(b.strings, s)
	return b
}

func (b *BatchOverlap) Result() (result []string) {
	if b.hasEmpty {
		return
	}

	for idx := 0; idx < len(b.strings); idx++ {
		if result == nil {
			result = b.strings[0]
		}
		result = StringsOverlap(result, b.strings[idx])
	}
	return
}

func StringsOverlap(as, bs []string) []string {
	ret := make([]string, 0)
	for _, a := range as {
		for _, b := range bs {
			if a == b {
				ret = append(ret, a)
			}
		}
	}
	return Uniq(ret)
}

func RemoveElementInSlice(strs []string, ele string) []string {
	if !Contains(strs, ele) {
		return strs
	}
	ret := []string{}
	for _, s := range strs {
		if ele == s {
			continue
		}
		ret = append(ret, s)
	}
	return ret
}

func SaveServerOut(stage *log.Stage, rootDir, filename string) {
	os.Rename(fmt.Sprintf("%v/bin/%v.1", rootDir, filename), fmt.Sprintf("%v/bin/%v.2", rootDir, filename))
	os.Rename(fmt.Sprintf("%v/bin/%v", rootDir, filename), fmt.Sprintf("%v/bin/%v.1", rootDir, filename))
}

func TimestampToLocalString(timestamp int64) string {
	return time.Unix(0, timestamp).Format(log.LogTimeStamp)
}

func KvArrToMap(keys, vals []string) map[string]string {
	ret := make(map[string]string)
	for idx, _ := range keys {
		ret[keys[idx]] = vals[idx]
	}
	return ret
}

const md5SumReadBufferSize = 65536

// MD5sum returns MD5 checksum of filename
func MD5sum(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	for buf, reader := make([]byte, md5SumReadBufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		hash.Write(buf[:n])
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}

type ErrTypeDifferentError struct {
	srcType string
	dstType string
}

func (e *ErrTypeDifferentError) Error() string {
	return fmt.Sprintf("src type<%v> is different from dst type<%v>", e.srcType, e.dstType)
}

func newTypeDifferentError(srcType, dstType string) *ErrTypeDifferentError {
	return &ErrTypeDifferentError{srcType, dstType}
}

func DeepCopyViaJsonEncoding(src, dst interface{}) error {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	if srcType != dstType {
		return newTypeDifferentError(srcType.String(), dstType.String())
	}

	data, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, dst); err != nil {
		return err
	}

	return nil
}

// GetSliceAsStringLineByLine formats a string slice to a string with line breaks
// and line numbers.
func GetSliceAsStringLineByLine(messages []string) string {
	var sb strings.Builder
	for i := range messages {
		sb.WriteString(fmt.Sprintf("\n%v. %s", i+1, messages[i]))
	}
	return sb.String()
}

// ReadPidFile reads a pid file and return its pid value
func ReadPidFile(filePath string) (int32, error) {
	pidBs, err := ioutil.ReadFile(filePath)
	if nil != err {
		return 0, fmt.Errorf("[ERROR] Failed to read pid file:%s", err.Error())
	}
	pid, err := strconv.ParseInt(string(pidBs), 10, 32)
	if nil != err {
		return 0, fmt.Errorf("[ERROR] Failed to parse pid bytes(%s), error: %s", string(pidBs), err.Error())
	}
	return int32(pid), nil
}

func GetRelativeCNFPath(umcSubPath string, template string) string {
	return fmt.Sprintf("./%s/%s", umcSubPath, template)
}
