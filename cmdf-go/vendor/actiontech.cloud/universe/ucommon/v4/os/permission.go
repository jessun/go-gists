package os

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	builtin_syscall "syscall"
	"unsafe"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/util"
)

func SetCaps(stage *log.Stage, path string, caps string) error {
	// /dev/null for sudo password
	if ret, retCode, err := Cmdf(stage, "SUDO setcap %v=+eip %v </dev/null", caps, path); nil != err {
		return err
	} else if 0 != retCode {
		return fmt.Errorf("setcap %v error(%v)", path, ret)
	}
	return nil
}

func addUser(stage *log.Stage, runUser string) error {

	if _, retCode, err := Cmdf(stage, "grep %v: /etc/passwd", runUser); nil != err {
		return err
	} else if 0 == retCode {
		return nil
	}

	if ret, _, err := Cmdf(stage, "SUDO groupadd -f ACTIONTECH"); nil != err {
		return fmt.Errorf("groupadd failed, err=%v, ret=%v", err, ret)
	}
	if _, err := Cmdf2(stage, "SUDO useradd -M -s /bin/bash -r -g ACTIONTECH %v", runUser); nil != err {
		return fmt.Errorf("useradd failed, err=%v", err)
	}
	return nil
}

func DelUserCmd(user string) string {
	return fmt.Sprintf("userdel %v ||  [ $? -eq 6 ]", user)
}

func DelGroupCmd(group string) string {
	return fmt.Sprintf("groupdel %v ||  [ $? -eq 6 ]", group)
}

// @ATTENTION this func must invoked by root
func UpdateLimits(stage *log.Stage, runUser string) error {
	limitsDir := "/etc/security/limits.d"
	if err := EnsureDir(stage, limitsDir, "", 0755); nil != err {
		return err
	}

	if isDir, _ := util.IsDir(limitsDir); !isDir {
		return fmt.Errorf("%v is not a directory can't update run user limit config", limitsDir)
	}

	limitPath := filepath.Join(limitsDir, fmt.Sprintf("actiontech.%v.conf", runUser))
	if IsFileExist(limitPath) {
		return nil
	}
	limitsCon := GetLimitsFileContents(stage, runUser)
	log.Brief(stage, "create %v, content: (%v)", limitPath, limitsCon)

	if err := SafeWriteFile(stage, limitPath, limitsCon, "", 0644); nil != err {
		return err
	}
	if err := os.Chmod(limitPath, 0644); nil != err {
		return err
	}
	return nil
}

func GetLimitsFileContents(stage *log.Stage, runUser string) string {
	// create user limits file
	rlimit, err := GetNumFromFile("/proc/sys/fs/nr_open")
	if nil != err {
		rlimit = 1048576
		log.Key(stage, err.Error())
	}
	return fmt.Sprintf(`%v soft nofile %v
%v hard nofile %v
%v soft nproc %v
%v hard nproc %v
%v hard sigpending %v
%v soft sigpending %v
`, runUser, rlimit, runUser, rlimit, runUser,
		rlimit, runUser, rlimit, runUser, rlimit, runUser, rlimit)
}

func getRootSudoUpdateContent() string {
	return "Defaults:root !requiretty\n"
}

func GetSudoUpdateContent(stage *log.Stage, runUser string) string {
	sudoerLines := []string{
		fmt.Sprintf("Defaults:%v !requiretty", runUser),
	}
	{
		cmds := []string{}
		for _, command := range []string{"mount", "umount", "lsof", "setcap", "sg_persist", "ip", "arping", "su",
			"useradd", "userdel", "ls", "groupadd", "groupdel", "chown", "tar", "cp", "chmod", "chkconfig",
			"update-rc.d", "stat", "multipath", "mpathpersist", "systemctl", "blkid"} {
			path, retCode, _ := Cmdf(stage, "which %v", command)
			if 0 != retCode {
				log.Brief(stage, "command %v is missing in this system", command)
				continue
			}
			cmds = append(cmds, path)
		}
		sudoerLines = append(sudoerLines, fmt.Sprintf("%v ALL=(ALL) NOPASSWD: %v", runUser, strings.Join(cmds, ",")))
	}
	return strings.Join(sudoerLines, "\n") + "\n"
}

func InstallUser(stage *log.Stage, runUser string) error {
	stage.Enter("Install_user")
	defer stage.Exit()

	if "root" == runUser {
		return UpdateSudoers(stage, runUser, getRootSudoUpdateContent())
	} else {
		if _, retCode, err := Cmdf(stage, "grep %v: /etc/passwd", runUser); nil != err {
			return err
		} else if 0 == retCode {
			return nil
		}

		if err := addUser(stage, runUser); nil != err {
			return err
		}
		if err := UpdateSudoers(stage, runUser, GetSudoUpdateContent(stage, runUser)); nil != err {
			return err
		}
		return nil
	}
}

func UpdateSudoers(stage *log.Stage, runUser, sudoCon string) error {
	//check includedir in sudoers
	if i, err := ioutil.ReadFile("/etc/sudoers"); nil != err {
		return fmt.Errorf("read /etc/sudoers failed, err=%v", err)
	} else {
		if 0 == len(regexp.MustCompile(`(?m)^\s*#includedir\s+/etc/sudoers\.d\s*$`).FindStringSubmatch(string(i))) {
			return fmt.Errorf("/etc/sudoers doesn't contain \"#includedir /etc/sudoers.d\" line")
		}
	}

	// create user sudoer file
	if !IsFileExist("/etc/sudoers.d") {
		return fmt.Errorf("/etc/sudoers.d doesn't exist, can't update run user sudo permission")
	}
	sudoersFile := fmt.Sprintf("/etc/sudoers.d/actiontech-%v", runUser)
	if err := SafeWriteFile(stage, sudoersFile, sudoCon, "root", 0440); nil != err {
		return err
	}
	return nil
}

func CheckSudo(stage *log.Stage, runUser string) error {
	if "root" == runUser {
		return nil
	}
	if _, err := Cmdf2(stage, "su %v -c \"sudo su root -c 'ls'\"", runUser); nil != err {
		return fmt.Errorf("sudo error: %v", err)
	}
	return nil
}

func SetKeepCaps() error {
	_, _, err := builtin_syscall.RawSyscall(builtin_syscall.SYS_PRCTL, builtin_syscall.PR_SET_KEEPCAPS, 1, 0)
	if 0 != err {
		return err
	}
	return nil
}

// gets the capabilities from process pid
func GetCap(pid int) (*Cap, error) {
	caps := new(Cap)
	caps.Header.Version = LINUX_CAPABILITY_VERSION_3
	caps.Header.Pid = int32(pid)
	_, _, e1 := builtin_syscall.Syscall(builtin_syscall.SYS_CAPGET,
		uintptr(unsafe.Pointer(&caps.Header)),
		uintptr(unsafe.Pointer(&caps.Data[0])),
		0)
	if e1 != 0 {
		return nil, e1
	}
	return caps, nil
}

// set the capabilities
func SetCap(capData CapData) error {
	caps := new(Cap)
	caps.Header.Version = LINUX_CAPABILITY_VERSION_3
	caps.Data[0] = capData
	_, _, e1 := builtin_syscall.Syscall(builtin_syscall.SYS_CAPSET,
		uintptr(unsafe.Pointer(&caps.Header)),
		uintptr(unsafe.Pointer(&caps.Data[0])),
		0)
	if e1 != 0 {
		return e1
	}
	return nil
}
