package nsnet

import (
	"fmt"
	"net"
	"time"
)

// network interface check
// waits for netsetgo to creates tunneling between host and container

func WaitForNetwork() error {
	maxAttempt := 3
	checkInterval := time.Second
	for i := 0; i < maxAttempt; i++ {
		interfaces, err := net.Interfaces()
		if err != nil {
			return err
		}
		if len(interfaces) > 1 {
			return nil
		}
		time.Sleep(checkInterval)
	}
	return fmt.Errorf("too much attempts of waiting for network")
}
