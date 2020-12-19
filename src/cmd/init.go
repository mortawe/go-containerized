package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/mortawe/go-containerized/src/nsexec"
	"github.com/mortawe/go-containerized/src/nsnet"
	"github.com/mortawe/go-containerized/src/nsrootfs"
)

func init() {
	nsexec.Register("nsInit", nsInit)
	if nsexec.Init() {
		os.Exit(0)
	}
}

func nsInit() {
	newRootPath := os.Args[1]
	hostname := os.Args[2]

	if err := nsrootfs.MountProc(newRootPath); err != nil {
		fmt.Printf("Error mounting /proc - %s, newroot - %s\n", err, newRootPath)
		os.Exit(1)
	}

	if err := nsrootfs.PivotRoot(newRootPath); err != nil {
		fmt.Printf("Error privot root - %s, newroot - %s\n", err, newRootPath)
		os.Exit(1)
	}

	if err := syscall.Sethostname([]byte(hostname)); err != nil {
		fmt.Printf("Error setting hostname - %s, hostname - %s\n", err, hostname)
		os.Exit(1)
	}

	if err := nsnet.WaitForNetwork(); err != nil {
		fmt.Printf("Error waiting for network - %s\n", err)
		os.Exit(1)
	}

	run()
}
