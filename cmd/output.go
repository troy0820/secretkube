package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/kubernetes/pkg/apis/core"
	"os"
)

func printError(err error, cmd *cobra.Command, msg string) {
	if err != nil {
		cmd.Println(msg, err.Error())
	}
}

func printErrorWithExit(err error, cmd *cobra.Command, msg string) {
	if err != nil {
		cmd.Println(msg, err.Error())
		os.Exit(1)
	}
}

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Creates output of the secret",
	Long: `Creates the output of the secret you want to
create.  This output can be saved to a file or printed to the screen`,
	Run: func(cmd *cobra.Command, args []string) {
		fl, err := cmd.Flags().GetString("file")
		printError(err, cmd, "Error:")

		m, err := makeMapfromJson(fl)
		printErrorWithExit(err, cmd, "Error:")
		cmd.Println("map", m)

		out, err := cmd.Flags().GetString("output")
		printError(err, cmd, "Error:")

		ns, err := cmd.Flags().GetString("namespace")
		printError(err, cmd, "Error:")

		clientset := fake.NewSimpleClientset()
		cmd.Println("clientset", clientset)

		objMeta := metav1.ObjectMeta{
			Name: "fake-secret",
		}
		objTypeMeta := metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		}
		cmd.Println("==============================================")
		bytemap := convertMapValuesToBase64(turnMaptoBytes(m))
		sec := core.Secret{TypeMeta: objTypeMeta, ObjectMeta: objMeta, Data: bytemap}
		clientset.CoreV1().Secrets(ns).Create(&v1.Secret{
			ObjectMeta: objMeta,
			Data:       bytemap,
		})
		secrets, err := clientset.CoreV1().Secrets("").List(metav1.ListOptions{})
		printError(err, cmd, "Error:")
		cmd.Println("These are the secrets", secrets)
		cmd.Println("===============================================")
		if len(secrets.Items) > 0 {
			for _, secret := range secrets.Items {
				cmd.Println("this is the for loop", secret.GetName(), secret.GetNamespace())
			}
		}
		cmd.Printf("secret: %+v\n\n", sec)
		cmd.Printf("meta: %+v\n\n", objMeta)
		cmd.Printf("Type: %+v\n\n", objTypeMeta)
		if fl != "" && out != "" && ns != "" {
			cmd.Printf("Saving %s secret to: %s in %s namespace", convertToBase64(fl), out, ns)
		} else {
			cmd.Println("No file location chosen")
			os.Exit(1)
		}
	},
}
