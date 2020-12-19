package nsexec

import (
	"fmt"
	"os"
)

var registeredInitializers = make(map[string]func())

func Register(name string, initializer func()) error {
	if _, exists := registeredInitializers[name]; exists {
		panic(fmt.Sprintf("reexec func already registered under name %q", name))
	}
	registeredInitializers[name] = initializer
	return nil
}

func Init() bool {
	initializer, exists := registeredInitializers[os.Args[0]]
	if exists {
		initializer()
		return true
	}
	return false
}
