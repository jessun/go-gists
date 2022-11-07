//+build linux

package os

import (
	"fmt"
	os_ "os"
	"runtime"

	"golang.org/x/sys/unix"

	"actiontech.cloud/universe/ucommon/v4/log"
	user_ "actiontech.cloud/universe/ucommon/v4/user"
	syscall "actiontech.cloud/universe/ucommon/v4/wrapper/syscall"
)

const (
	ACCESS_W = unix.W_OK
	ACCESS_R = unix.R_OK
	ACCESS_X = unix.X_OK
)

func CheckAccess(user string, path string, requiredAccess uint32) error {
	stage := log.NewStage().Enter(fmt.Sprintf(
		"check_user[%v]_access[%v]_to_path[%v]", user, requiredAccess, path))
	defer stage.Exit()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	originalUid := os_.Getuid()
	originalGid := os_.Getgid()

	var uid int
	var gids []int

	var err error
	uid, gids, err = user_.LookupUidGidByUser(user)
	if nil != err {
		return err
	}

	var accessErr error

	for _, gid := range gids {
		originalCap, err := GetCap(0)
		if nil != err {
			return fmt.Errorf("GetCap error: %v", err)
		}

		//change user, and clear CAP
		{
			if err := SetKeepCaps(); nil != err {
				return fmt.Errorf("SetKeepCaps error: %v", err)
			}

			if err := syscall.Setresgid(stage, gid, gid, gid); nil != err {
				return fmt.Errorf("Setresgid error: %v", err)
			}

			if err := syscall.Setresuid(stage, uid, uid, uid); nil != err {
				return fmt.Errorf("Setresuid error: %v", err)
			}

			capData := CapData{
				Effective:   EncodeCaps([]Capability{CAP_SETGID, CAP_SETUID}),
				Permitted:   originalCap.EncodePermittedCaps(),
				Inheritable: originalCap.EncodeInheritableCaps(),
			}

			if err := SetCap(capData); nil != err {
				return fmt.Errorf("SetCap error: %v", err)
			}
		}

		accessErr = unix.Access(path, requiredAccess)

		//change user back, and clear
		{
			if err := SetKeepCaps(); nil != err {
				return fmt.Errorf("SetKeepCaps error: %v", err)
			}

			if err := syscall.Setresgid(stage, originalGid, originalGid, originalGid); nil != err {
				return fmt.Errorf("Setresgid as original error: %v", err)
			}

			if err := syscall.Setresuid(stage, originalUid, originalUid, originalUid); nil != err {
				return fmt.Errorf("Setresuid as original error: %v", err)
			}

			if err := SetCap(originalCap.Data[0]); nil != err {
				return fmt.Errorf("SetCap as original error: %v", err)
			}
		}

		if nil == accessErr {
			return nil
		}
	}

	return accessErr
}
