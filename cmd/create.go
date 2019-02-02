package cmd

import (
	"encoding/base64"
	"github.com/spf13/cobra"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/fatih/color"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func convertToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//TODO: use createSecret with input from file to create secret
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates secret for cluster",
	Long:    `Creates a kubernetes secret for your cluster`,
	Aliases: []string{"Create"},
	Run: func(cmd *cobra.Command, args []string) {

		var kubeconfig, ns string
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
		if cmd.Flags().Changed("config") {
			kubeconfig, _ = cmd.Flags().GetString("config")
		}
		ns = "default"
		if cmd.Flags().Changed("namespace") {
			ns, _ = cmd.Flags().GetString("namespace")
		}
		file, _ := cmd.Flags().GetString("file")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		printErrorWithExit(err, cmd, "Error:")
		cmd.Println(green("Kubeconfig:"), green(config))
		clientset := kubernetes.NewForConfigOrDie(config)
		//		secretclient := clientset.CoreV1().Secrets(ns)
		cmd.Println(green("Kubernetes clientset"), red(clientset))
		printError(err, cmd, "Error:")

		// TODO:// Take json file and use it to create secret to pass into client
		/*	m, err := makeMapfromJson(fl)
			printErrorWithExit(err, cmd, "Error:")
			stringdata := turnMaptoString(m)
			bytemap := turnMaptoBytes(m)
			convertMapValuesToBase64(bytemap)
			sec, err := createSecret(name, stringdata, bytemap)
			printError(err, cmd, "Error:")
			secretclient.Create(sec)
		*/

		cmd.Println(green("This is a string base64 encoded"), red(file), red(ns), config)
	},
}
