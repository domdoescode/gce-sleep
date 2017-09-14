package cmd

import (
	"log"
)

var (
	verbose        bool
	configLocation string
	labelName      string

	version string
	commit  string
	date    string
)

// Setup passes through ldflags from goreleaser, adds child commands, and sets up flag configuration
func Setup(v, c, d string) {
	version = v
	commit = c
	date = d

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "toggles verbose output")
	RootCmd.PersistentFlags().StringVarP(&configLocation, "config", "c", "/etc/gce-sleep.conf", "gce-sleep config file")
	RootCmd.PersistentFlags().StringVarP(&labelName, "label", "l", "gce-sleep", "label for filtering instances")

	RootCmd.AddCommand(versionCmd)
}

func logPrintlnVerbose(v ...interface{}) {
	if verbose {
		log.Println(v...)
	}
}

func logPrintfVerbose(format string, v ...interface{}) {
	if verbose {
		log.Printf(format, v...)
	}
}
