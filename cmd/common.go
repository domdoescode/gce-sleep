package cmd

var (
	verbose        bool
	configLocation string
	labelName      string

	version string
	tag     string
	commit  string
	date    string
)

func Setup(v, t, c, d string) {
	version = v
	tag = t
	commit = c
	date = d

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "toggles verbose output")
	RootCmd.PersistentFlags().StringVarP(&configLocation, "config", "c", "/etc/gce-sleep.conf", "gce-sleep config file")
	RootCmd.PersistentFlags().StringVarP(&labelName, "label", "l", "gce-sleep", "label for filtering instances")
}
