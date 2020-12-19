package nsexec

import (
	"os/exec"
	"syscall"
)

// symbolic link for runtime file in current env
func self() string {
	return "/proc/self/exe"
}

func Command(args ...string) *exec.Cmd {
	return &exec.Cmd{
		Path: self(),
		Args: args,
		SysProcAttr: &syscall.SysProcAttr{
			Pdeathsig: syscall.SIGTERM,
		},
	}
}
