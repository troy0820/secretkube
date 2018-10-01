package cmd

import (
	"encoding/base64"
	"github.com/spf13/cobra"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func convertToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates secret for cluster",
	Long:    `Creates a kubernetes secret for your cluster`,
	Aliases: []string{"Create"},
	Run: func(cmd *cobra.Command, args []string) {

		var kubeconfig, str, namespace string
		fl, err := cmd.Flags().GetString("config")
		printError(err, cmd, "Error:")
		if fl != "" {
			str = "Flag is set: " + convertToBase64(fl)
			kubeconfig = fl
			cmd.Println("Kubeconfig: ", kubeconfig)

		} else {
			str = "Flag is not set: default: " + convertToBase64("This is the default string")
			kubeconfig = os.Getenv("HOME") + "/.kube/config"
			cmd.Println("Kubeconfig: ", kubeconfig)
		}

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		printError(err, cmd, "Error:")
		cmd.Println("Kubeconfig:", config)
		clientset := kubernetes.NewForConfigOrDie(config)
		cmd.Println("Kubernetes clientset", clientset)
		ns, err := cmd.Flags().GetString("namespace")
		printError(err, cmd, "Error:")

		if ns != "" {
			namespace = "Hello"
		} else {
			namespace = ""
		}

		cmd.Println("This is a string base64 encoded", str, namespace)
	},
}
