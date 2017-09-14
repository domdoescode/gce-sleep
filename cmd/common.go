package cmd

var (
	verbose                   bool
	configLocation, labelName string
)

func Setup() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "toggles verbose output")
	RootCmd.PersistentFlags().StringVarP(&configLocation, "config", "c", "/etc/gce-sleep.conf", "gce-sleep config file")
	RootCmd.PersistentFlags().StringVarP(&labelName, "label", "l", "gce-sleep", "label for filtering instances")
}
