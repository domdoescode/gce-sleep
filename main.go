package main

import (
	"log"

	"github.com/domudall/gce-sleep/cmd"
)

var (
	version = "master"
	tag     = "tag"
	commit  = "commit"
	date    = "date"
)

func main() {
	cmd.Setup(version, tag, commit, date)
	cmd.RootCmd.AddCommand(cmd.VersionCmd)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
