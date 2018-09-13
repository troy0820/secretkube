package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionNumber = "0.0.1"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints the version of secretkube",
	Aliases: []string{"Version"},
	Run: func(cmd *cobra.Command, args []string) {
		Version(versionNumber)
	},
}

func Version(str string) {
	fmt.Println("SecretKube -- version", str)
}
