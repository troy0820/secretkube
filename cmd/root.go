package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "secretkube",
	Short: "Create Kubernetes secrets with key/value pairs from json",
	Long: `Allows you to create Kubernetes secrets with json from key/value pairs
		and place them in your cluster without creating yaml files`,
	Run: func(cmd *cobra.Command, args []string) {
		fl, err := cmd.Flags().GetBool("version")
		if err != nil {
			cmd.Println(err.Error())
		}
		if fl {
			Version()
		} else {
			cmd.Help()
		}
	},
}

var vers bool

func init() {
	rootCmd.AddCommand(versionCmd())
	rootCmd.Flags().BoolVarP(&vers, "version", "v", false, "version output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
