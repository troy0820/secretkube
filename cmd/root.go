package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "secretkube",
	Short: "Create Kubernetes secrets with key/value pairs from json",
	Long: `Create Kubernetes secrets with json from key/value pairs
and place them in your cluster without creating yaml files`,
	Run: func(cmd *cobra.Command, args []string) {
		fl, err := cmd.Flags().GetBool("version")
		if err != nil {
			cmd.Println(err.Error())
		}
		if fl {
			cmd.Println(Version(versionNumber))
		} else {
			cmd.Help()
		}
	},
}

var vers bool
var config, namespace, jsonFile, file, output, ns string

func init() {
	createCmd.Flags().StringVarP(&config, "config", "c", "", "filepath of kubeconfig")
	createCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace to put secret in")
	createCmd.Flags().StringVarP(&jsonFile, "file", "f", "", "filepath of JSON file")
	outputCmd.Flags().StringVarP(&file, "file", "f", "", "file path of JSON file")
	outputCmd.Flags().StringVarP(&output, "output", "o", "", "output file to save secret")
	outputCmd.Flags().StringVarP(&ns, "namespace", "n", "", "namespace for secret")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(outputCmd)
	rootCmd.Flags().BoolVarP(&vers, "version", "v", false, "version output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
