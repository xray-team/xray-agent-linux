package conf

import "flag"

// Flags defines configuration passed by flags.
type Flags struct {
	ConfigFilePath *string
	DryRun         *bool
}

func ParseFlags() *Flags {
	var f Flags

	f.ConfigFilePath = flag.String("config", "./config.json", "path to config file")
	f.DryRun = flag.Bool("dryrun", false, "test run")

	flag.Parse()

	return &f
}
