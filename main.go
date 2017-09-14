package main

import (
	"log"

	"github.com/domudall/gce-sleep/cmd"
)

func main() {
	cmd.Setup()
	cmd.RootCmd.AddCommand(cmd.VersionCmd)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
