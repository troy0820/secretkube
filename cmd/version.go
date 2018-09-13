package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Prints the version of secretkube",
		Aliases: []string{"Version"},
		Run: func(cmd *cobra.Command, args []string) {
			Version()
		},
	}
}

func Version() {
	fmt.Println("SecretKube -- version 0.0.1")
}
