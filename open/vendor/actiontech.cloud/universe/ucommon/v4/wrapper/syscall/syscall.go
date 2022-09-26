/*
This file wraps all built-in syscall functions used in DMP,
for more info, see: http://10.186.18.11/jira/browse/DMP-4106
*/

package syscall

import (
	"fmt"
	"strings"
	builtin_syscall "syscall"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/wrapper"
)

// wrapper for syscall.Kill. descriptions simply describes why to call this func(empty is allowed), eg: tgt=umc,destn=dump
func Kill(stage *log.Stage, pid int, sig builtin_syscall.Signal, descriptions ...string) (err error) {
	from := wrapper.GetCurrentTimestamp()
	err = builtin_syscall.Kill(pid, sig)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)kill_pid[%v]sig[%v]", from, to, pid, sig))
	defer stage.Exit()
	desc := strings.Join(descriptions, ",")

	if nil != err {
		log.Write(stage).
			Brief("err:[%v]", err.Error()).
			Detail(desc).
			Done()
	} else {
		log.Detail(stage, "succeeded.%v", desc)
	}
	return
}

// Setresgid is a wrapper for RawSyscall. descriptions simply describes why to call this func(empty is allowed), eg: destn=clear_gid
func Setresgid(stage *log.Stage, rgid, egid, sgid int, descriptions ...string) error {
	from := wrapper.GetCurrentTimestamp()
	// After Go1.16, calling Setresuid/Setresgid affects all threads of the current process.
	// We need call RawSyscall directly for keeping the same behavior with before Go1.16.
	//
	// ref: https://github.com/golang/go/issues/1435, http://10.186.18.11/jira/browse/DMP-11793
	_, _, errno := builtin_syscall.RawSyscall(builtin_syscall.SYS_SETRESGID, uintptr(rgid), uintptr(egid), uintptr(sgid))
	//err = builtin_syscall.Setresgid(rgid, egid, sgid)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)set_rgid[%v]egid[%v]sgid[%v]", from, to, rgid, egid, sgid))
	defer stage.Exit()
	desc := strings.Join(descriptions, ",")

	if errno != 0 {
		log.Write(stage).
			Brief("err:[%v]", errno.Error()).
			Detail(desc).
			Done()
		return errno
	}
	log.Detail(stage, "succeeded.%v", desc)
	return nil
}

// Setresuid is a wrapper for RawSyscall. descriptions simply describes why to call this func(empty is allowed), eg: destn=clear_uid
func Setresuid(stage *log.Stage, ruid, euid, suid int, descriptions ...string) error {
	from := wrapper.GetCurrentTimestamp()
	// After Go1.16, calling Setresuid/Setresgid affects all threads of the current process.
	// We need call RawSyscall directly for keeping the same behavior with before Go1.16.
	//
	// ref: https://github.com/golang/go/issues/1435, http://10.186.18.11/jira/browse/DMP-11793
	_, _, errno := builtin_syscall.RawSyscall(builtin_syscall.SYS_SETRESUID, uintptr(ruid), uintptr(euid), uintptr(suid))
	//err = builtin_syscall.Setresuid(ruid, euid, suid)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)set_ruid[%v]euid[%v]suid[%v]", from, to, ruid, euid, suid))
	defer stage.Exit()
	desc := strings.Join(descriptions, ",")

	if errno != 0 {
		log.Write(stage).
			Brief("err:[%v]", errno.Error()).
			Detail(desc).
			Done()
		return errno
	}
	log.Detail(stage, "succeeded.%v", desc)
	return nil
}

// wrapper for syscall.Getpid. descriptions simply describes why to call this func(empty is allowed), eg: tgt=self
func Getpid(stage *log.Stage, descriptions ...string) (pid int) {
	from := wrapper.GetCurrentTimestamp()
	pid = builtin_syscall.Getpid()
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)get_pid[%v]", from, to, pid))
	defer stage.Exit()
	desc := strings.Join(descriptions, ",")

	log.Detail(stage, "succeeded.%v", desc)

	return
}

// wrapper for syscall.Getrlimit. descriptions simply describes why to call this func(empty is allowed), eg: tgt=self
func Getrlimit(stage *log.Stage, resource int, rlim *builtin_syscall.Rlimit, descriptions ...string) (err error) {
	from := wrapper.GetCurrentTimestamp()
	err = builtin_syscall.Getrlimit(resource, rlim)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)get_rlimit: resource[%v]rlim[%#v]", from, to, resource, rlim))
	defer stage.Exit()
	desc := strings.Join(descriptions, ",")

	if err != nil {
		log.Write(stage).
			Brief("err:[%v]", err.Error()).
			Detail(desc).
			Done()
	} else {
		log.Detail(stage, "succeeded.%v", desc)
	}

	return
}

// wrapper for syscall.Setrlimit. descriptions simply describes why to call this func(empty is allowed), eg: tgt=self
func Setrlimit(stage *log.Stage, resource int, rlim *builtin_syscall.Rlimit, descriptions ...string) (err error) {
	from := wrapper.GetCurrentTimestamp()
	err = builtin_syscall.Setrlimit(resource, rlim)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)set_rlimit: resource[%v]rlim[%#v]", from, to, resource, rlim))
	defer stage.Exit()
	desc := strings.Join(descriptions, ",")

	if err != nil {
		log.Write(stage).
			Brief("err:[%v]", err.Error()).
			Detail(desc).
			Done()
	} else {
		log.Detail(stage, "succeeded.%v", desc)
	}

	return
}

// wrapper for syscall.Flock. descriptions simply describes why to call this func(empty is allowed), eg: tgt=self
func Flock(stage *log.Stage, fd, how int, descriptions ...string) (err error) {
	from := wrapper.GetCurrentTimestamp()
	err = builtin_syscall.Flock(fd, how)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)flock: fd[%v]how[%v]", from, to, fd, how))
	defer stage.Exit()

	desc := strings.Join(descriptions, ",")
	if err != nil {
		log.Write(stage).
			Brief("err:[%v]", err.Error()).
			Detail(desc).
			Done()
	} else {
		log.Detail(stage, "succeeded.%v", desc)
	}

	return
}

// wrapper for syscall.Statfs. descriptions simply describes why to call this func(empty is allowed), eg: tgt=self
func Statfs(stage *log.Stage, path string, buf *builtin_syscall.Statfs_t, descriptions ...string) (err error) {
	from := wrapper.GetCurrentTimestamp()
	err = builtin_syscall.Statfs(path, buf)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)statfs_path[%v]buf[%#v]", from, to, path, buf))
	defer stage.Exit()

	desc := strings.Join(descriptions, ",")
	if err != nil {
		log.Write(stage).
			Brief("err:[%v]", err.Error()).
			Detail(desc).
			Done()
	} else {
		log.Detail(stage, "succeeded.%v", desc)
	}

	return
}

// wrapper for syscall.Stat. descriptions simply describes why to call this func(empty is allowed), eg: tgt=self
func Stat(stage *log.Stage, path string, stat *builtin_syscall.Stat_t, descriptions ...string) (err error) {
	from := wrapper.GetCurrentTimestamp()
	err = builtin_syscall.Stat(path, stat)
	to := wrapper.GetCurrentTimestamp()

	stage.Enter(fmt.Sprintf("(%v~%v)state_path[%v]stat[%#v]", from, to, path, stat))
	defer stage.Exit()

	desc := strings.Join(descriptions, ",")
	if err != nil {
		log.Write(stage).
			Brief("err:[%v]", err.Error()).
			Detail(desc).
			Done()
	} else {
		log.Detail(stage, "succeeded.%v", desc)
	}
	return
}
