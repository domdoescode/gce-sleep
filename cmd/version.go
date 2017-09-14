package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gce-sleep",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.1")
	},
}
