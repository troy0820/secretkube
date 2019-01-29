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
//TODO: use cmd.Flags().Changed('string') to gather flags for create command
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Creates secret for cluster",
	Long:    `Creates a kubernetes secret for your cluster`,
	Aliases: []string{"Create"},
	Run: func(cmd *cobra.Command, args []string) {

		var kubeconfig, str, namespace string
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
		if cmd.Flags().Changed("config") {
			kubeconfig, _ = cmd.Flags().GetString("config")
		}
		cmd.Println(green("Using Kubeconfig: ", green(kubeconfig)))

		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		printErrorWithExit(err, cmd, "Error:")
		cmd.Println("Kubeconfig:", config)
		clientset := kubernetes.NewForConfigOrDie(config)
		cmd.Println("Kubernetes clientset", clientset)
		ns, err := cmd.Flags().GetString("namespace")
		printError(err, cmd, "Error:")
		/* TODO:// Take json file and use it to create secret to pass into client
		m, err := makeMapfromJson(fl)
		printErrorWithExit(err, cmd, "Error:")
		stringdata := turnMaptoString(m)
		bytemap := turnMaptoBytes(m)
		convertMapValuesToBase64(bytemap)
		sec, err := createSecret(name, stringdata, bytemap)
		printError(err, cmd, "Error:")
		*/
		if ns != "" {
			namespace = "Hello"
		} else {
			namespace = ""
		}

		cmd.Println("This is a string base64 encoded", str, namespace)
	},
}
