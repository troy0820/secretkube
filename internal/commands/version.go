package commands

import (
	"github.com/spf13/cobra"
	"github.com/troy0820/secretkube/version"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints the version of secretkube",
	Aliases: []string{"Version"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(Version(version.Version))
	},
}

//Version  command takes the string and shows the version
func Version(str string) string {
	return "SecretKube -- version " + str
}
