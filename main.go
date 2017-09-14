package main

import (
	"log"

	"github.com/domudall/gce-sleep/cmd"
)

var (
	version = "master"
	commit  = "commit"
	date    = "date"
)

func main() {
	cmd.Setup(version, commit, date)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
