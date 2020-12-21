package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/mortawe/go-containerized/src/nsexec"
	"github.com/mortawe/go-containerized/src/nsnet"
	"github.com/mortawe/go-containerized/src/nsopts"
)

func main() {
	opts := nsopts.NewOpts()
	if !opts.Validate() {
		opts.Help()
		os.Exit(1)
	}
	// invoke self exec to isolate hostname and fs from host system
	cmd := nsexec.Command("nsInit", opts.GetRootfs(), opts.GetHostname())

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error invoking self exec - %s\n", err)
		os.Exit(1)
	}

	pid := fmt.Sprintf("%d", cmd.Process.Pid)
	// creating network bridge and veths between host system and container
	if err := nsnet.InvokeNetsetgo(opts.GetNetsetgo(), pid); err != nil {
		fmt.Printf("Error running netsetgo - %s\n", err)
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for invoked self - %s\n", err)
		os.Exit(1)
	}
}
