package main

import (
	"fmt"
	"os"
	"os/exec"
)

func run() {
	// shell starting inside of container
	cmd := exec.Command("/bin/sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// colourful prompt string formatting =)
	cmd.Env = []string{`PS1=🦈 \e[0;95;1m$(whoami)@$(hostname)\e[m\:\e[0;33;1m~$(pwd)\e[m$ `}

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		os.Exit(1)
	}
}
