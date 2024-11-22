package flags

import (
	"flag"
)

var (
	ConfigPath string
)

func init() {
	flag.StringVar(&ConfigPath, "config-path",
		"configs/serviceconfig.toml",
		"path to config file")

	flag.Parse()
}
