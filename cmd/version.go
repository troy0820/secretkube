package cmd

import (
	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Prints the version of secretkube",
		Aliases: []string{"Version"},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("SecretKube -- version 0.0.1")
		},
	}
}
