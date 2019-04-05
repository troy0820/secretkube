package cmd

import (
	"encoding/base64"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" //Needed for side effects for GCP
	"k8s.io/client-go/tools/clientcmd"
)

func convertToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates secret for cluster",
	Long:    `Creates a kubernetes secret for your cluster`,
	Aliases: []string{"Create"},
	Run: func(cmd *cobra.Command, args []string) {

		var kubeconfig, ns string
		name, err := cmd.Flags().GetString("name")
		printError(err, cmd, "Error: ")
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
		if cmd.Flags().Changed("config") {
			var err error
			kubeconfig, err = cmd.Flags().GetString("config")
			printError(err, cmd, "Error:")
		}
		ns = "default"
		if cmd.Flags().Changed("namespace") {
			var err error
			ns, err = cmd.Flags().GetString("namespace")
			printError(err, cmd, "Error:")
		}
		file, err := cmd.Flags().GetString("file")
		printError(err, cmd, "Error:")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		printErrorWithExit(err, cmd, "Error:")
		clientset := kubernetes.NewForConfigOrDie(config)
		secretclient := clientset.CoreV1().Secrets(ns)
		//Create map from JSON file
		m, err := MakeMapFromJSON(file)
		printErrorWithExit(err, cmd, "Error: ")
		byteData := TurnMapToBytes(m)
		//Create secret with byteData
		sec, err := CreateSecret(name, byteData)
		printError(err, cmd, "Error")
		cmd.Println(green("Creating secret ", name))
		secret, err := secretclient.Create(sec)
		printError(err, cmd, red("Error: "))
		cmd.Printf("Secret %s created: \n", secret.Name)
	},
}
