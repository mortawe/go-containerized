package nsopts

import (
	"flag"
	"os"
)

type Opts struct {
	hostname string
	rootfs   string
	netsetgo string
}

func NewOpts() Opts {
	o := Opts{}
	flag.StringVar(&o.hostname, "hostname", "go-containerized", "hostname inside container")
	flag.StringVar(&o.rootfs, "rootfs", "/tmp/go-containerized/rootfs", "path to the root filesystem to use")
	flag.StringVar(&o.netsetgo, "netsetgo", "/usr/local/bin/netsetgo", "path to the netsetgo binary")
	flag.Parse()
	return o
}

func (o Opts) Validate() bool {
	return exists(o.rootfs) && exists(o.netsetgo)
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func (o Opts) GetRootfs() string {
	return o.rootfs
}

func (o Opts) GetNetsetgo() string {
	return o.netsetgo
}

func (o Opts) GetHostname() string {
	return o.hostname
}
