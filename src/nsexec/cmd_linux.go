package nsexec

import (
	"os/exec"
	"syscall"
)

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