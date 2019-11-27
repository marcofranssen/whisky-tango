package main

import (
	"github.com/marcofranssen/whisky-tango/cmd/wt/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	v := cmd.VersionInfo{
		Version: version,
		Commit:  commit,
		Date:    cmd.ParseDate(date),
	}

	cmd.Execute(v)
}
