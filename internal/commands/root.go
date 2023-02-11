package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/troy0820/secretkube/version"
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
			cmd.Println(Version(version.Version))
		} else {
			cmd.Help()
		}
	},
}

var vers bool
var config, namespace, jsonFile, file, output, ns, name, sname, tname string

func init() {
	createCmd.Flags().StringVarP(&config, "config", "c", "", "filepath of kubeconfig")
	createCmd.Flags().StringVarP(&sname, "name", "n", "", "name of the secret")
	createCmd.Flags().StringVarP(&namespace, "namespace", "s", "", "namespace to put secret in")
	createCmd.Flags().StringVarP(&jsonFile, "file", "f", "", "filepath of JSON file")
	createCmd.MarkFlagRequired("file")
	createCmd.MarkFlagRequired("name")
	outputCmd.Flags().StringVarP(&file, "file", "f", "", "file path of JSON file")
	outputCmd.Flags().StringVarP(&output, "output", "o", "", "output file to save secret")
	outputCmd.Flags().StringVarP(&ns, "namespace", "s", "", "namespace for secret")
	outputCmd.Flags().StringVarP(&name, "name", "n", "", "name of the secret")
	outputCmd.MarkFlagRequired("name")
	outputCmd.MarkFlagRequired("file")
	get.Flags().StringVarP(&tname, "name", "n", "", "name of the secret")
	watcherCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace of secrets to watch")
	rootCmd.AddCommand(versionCmd, createCmd, outputCmd, get, watcherCmd, informerCmd)
	rootCmd.Flags().BoolVarP(&vers, "version", "v", false, "version output")
}

//Execute command to execute secretkube command line
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
