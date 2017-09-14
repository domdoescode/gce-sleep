package main

import (
	"log"

	"github.com/domudall/gce-sleep/cmd"
)

var version = "master"

func main() {
	cmd.Setup(version)
	cmd.RootCmd.AddCommand(cmd.VersionCmd)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
