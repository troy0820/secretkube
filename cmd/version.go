package cmd

import "github.com/spf13/cobra"

var versionNumber = "0.0.1"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints the version of secretkube",
	Aliases: []string{"Version"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(Version(versionNumber))
	},
}

func Version(str string) string {
	return "SecretKube -- version " + str
}
