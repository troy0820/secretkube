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
		cmd.Println("This is a test command")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
