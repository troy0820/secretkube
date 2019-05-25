package cmd

import (
	"encoding/base64"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" //Needed for side effects for GCP
	"k8s.io/client-go/tools/clientcmd"
)

func convertToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

//CreateSecret creates a secret to put into cluster
func CreateSecret(name string, data map[string][]byte) (*v1.Secret, error) {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		Data: data,
		Type: "Opaque",
	}, nil

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
