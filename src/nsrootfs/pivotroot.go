package nsrootfs

import (
	"os"
	"path/filepath"
	"syscall"
)

// attaching new fs to container
func PivotRoot(newroot string) error {
	putOld := filepath.Join(newroot, "/.pivot_root")

	if err := syscall.Mount(
		newroot,
		newroot,
		"",
		syscall.MS_BIND|syscall.MS_REC,
		"",
	); err != nil {
		return err
	}

	if err := os.MkdirAll(putOld, 0700); err != nil {
		return err
	}

	if err := syscall.PivotRoot(newroot, putOld); err != nil {
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	putOld = "/.pivot_root"
	if err := syscall.Unmount(
		putOld,
		syscall.MNT_DETACH,
	); err != nil {
		return err
	}

	if err := os.RemoveAll(putOld); err != nil {
		return err
	}

	return nil
}
