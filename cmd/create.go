package cmd

import (
	"encoding/base64"
	"github.com/spf13/cobra"
	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"k8s.io/client-go/kubernetes"
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
		fl, err := cmd.Flags().GetString("config")
		if err != nil {
			cmd.Println(err.Error())
		}
		kubeconfig := os.Getenv("HOME") + "/.kube/config"
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			cmd.Println(err.Error())
		}
		cmd.Println("Kubeconfig:", config)
		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			cmd.Println(err.Error())
		}

		var str2, namespace string
		if ns != "" {
			namespace = "Hello"
		} else {
			namespace = ""
		}

		if fl != "" {
			str2 = convertToBase64(fl)
		} else {
			//default kubeconfig
			//TODO: assign variable to ~/.kube/config
			str2 = convertToBase64("This is the default string")
		}
		cmd.Println("This is a string base64 encoded", str2, namespace)
	},
}
