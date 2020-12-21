package nsrootfs

import (
	"os"
	"path/filepath"
	"syscall"
)

// mounting /proc dir into container fs
// without this mount PID namespace could not work properly
func MountProc(newRoot string) error {
	source := "proc"
	target := filepath.Join(newRoot, source)
	fstype := "proc"
	flags := 0
	data := ""

	os.MkdirAll(target, 0755)
	if err := syscall.Mount(
		source,
		target,
		fstype,
		uintptr(flags),
		data,
	); err != nil {
		return err
	}

	return nil
}
