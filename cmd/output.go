package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"os"
	"unicode"
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

func createOutputSecret(sec *v1.Secret) string {
	var a string
	for k, v := range sec.StringData {
		if unicode.IsDigit(rune(v[0])) || unicode.IsLetter(rune(v[0])) {
			a += fmt.Sprintf("  %s: %s\n", string(k[1:len(k)-1]), convertToBase64(string(v[0:len(v)-1])))
		} else {
			a += fmt.Sprintf("  %s: %s\n", string(k[1:len(k)-1]), convertToBase64(string(v[1:len(v)-2])))
		}
	}
	secret := `
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  name: ` + sec.GetName() + `
  namespace: ` + sec.GetNamespace() + `
data:    
` + a
	return secret
}

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Creates output of the secret",
	Long: `Creates the output of the secret you want to
create.  This output can be saved to a file or printed to the screen`,
	Run: func(cmd *cobra.Command, args []string) {
		fl, err := cmd.Flags().GetString("file")
		printError(err, cmd, "Error:")

		name, err := cmd.Flags().GetString("name")
		printError(err, cmd, "Error:")

		m, err := makeMapfromJson(fl)
		printErrorWithExit(err, cmd, "Error:")

		out, err := cmd.Flags().GetString("output")
		printError(err, cmd, "Error:")

		var ns string
		if cmd.Flags().Changed("namespace") {
			ns, err = cmd.Flags().GetString("namespace")
			printError(err, cmd, "Error:")
		} else {
			ns = "default"
		}

		clientset := fake.NewSimpleClientset()

		objMeta := metav1.ObjectMeta{
			Name: name,
		}
		objTypeMeta := metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		}

		bytemap := convertMapValuesToBase64(turnMaptoBytes(m))

		secretclient := clientset.CoreV1().Secrets(ns)
		secretclient.Create(&v1.Secret{
			ObjectMeta: objMeta,
			TypeMeta:   objTypeMeta,
			Data:       bytemap,
			StringData: turnMaptoString(m),
			Type:       "Opaque",
		})

		secret, err := secretclient.Get(objMeta.GetName(), metav1.GetOptions{})
		printError(err, cmd, "Error:")

		saveToFile(createOutputSecret(secret), out)
		cmd.Printf("Secret saved to %s file \n", out)
		if fl != "" && out != "" && ns != "" && name != "" {
			cmd.Printf("Saving %s secret to: %s in %s namespace", convertToBase64(fl), out, ns)
		} else {
			cmd.Println("No file location chosen")
			os.Exit(1)
		}
	},
}
