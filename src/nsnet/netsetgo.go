package nsnet

import (
	"os/exec"
)

func InvokeNetsetgo(path string, pid string) error {
	netsetgoCmd := exec.Command(path, "-pid", pid)
	return netsetgoCmd.Run()
}
