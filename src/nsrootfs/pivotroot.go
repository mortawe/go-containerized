package nsrootfs

import (
	"os"
	"path/filepath"
	"syscall"
)

func PivotRoot(newroot string) error {
	putold := filepath.Join(newroot, "/.pivot_root")

	if err := syscall.Mount(
		newroot,
		newroot,
		"",
		syscall.MS_BIND|syscall.MS_REC,
		"",
	); err != nil {
		return err
	}

	if err := os.MkdirAll(putold, 0700); err != nil {
		return err
	}

	if err := syscall.PivotRoot(newroot, putold); err != nil {
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	putold = "/.pivot_root"
	if err := syscall.Unmount(
		putold,
		syscall.MNT_DETACH,
	); err != nil {
		return err
	}

	if err := os.RemoveAll(putold); err != nil {
		return err
	}

	return nil
}
