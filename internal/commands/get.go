package commands

import (
	"os"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var get = &cobra.Command{
	Use:     "get",
	Short:   "Gets secret for cluster",
	Long:    `Gets a kubernetes secret for your cluster`,
	Aliases: []string{"Create"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

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
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		printErrorWithExit(err, cmd, "Error:")
		clientset := kubernetes.NewForConfigOrDie(config)
		secretclient := clientset.CoreV1().Secrets(ns)
		//Create secret with byteData
		sec, err := secretclient.Get(ctx, name, metav1.GetOptions{})
		printError(err, cmd, red("Error: "))
		cmd.Printf("Secret %s created: \n", sec.Data["address"])
		cmd.Printf("Secret %s string data : \n", sec.StringData)
	},
}
