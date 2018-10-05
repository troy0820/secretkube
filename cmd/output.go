package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/kubernetes/pkg/apis/core"
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
		cmd.Println("map", m)
		out, err := cmd.Flags().GetString("output")
		printError(err, cmd, "Error:")

		ns, err := cmd.Flags().GetString("namespace")
		printError(err, cmd, "Error:")

		clientset := fake.NewSimpleClientset()
		cmd.Println("clientset", clientset)

		objMeta := metav1.ObjectMeta{
			Name: name,
		}
		objTypeMeta := metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		} //TODO: Take out ==== before release
		cmd.Println("==============================================")
		bytemap := convertMapValuesToBase64(turnMaptoBytes(m))
		sec := core.Secret{TypeMeta: objTypeMeta, ObjectMeta: objMeta, Data: bytemap}
		cmd.Println("sec:", sec)
		secretclient := clientset.CoreV1().Secrets(ns)
		secretclient.Create(&v1.Secret{
			ObjectMeta: objMeta,
			Data:       bytemap,
			StringData: turnMaptoString(m),
			Type:       "Opaque",
		})
		//TODO: Validate that secret created equals the secret made with createOutputSecret
		secret, err := secretclient.Get(objMeta.GetName(), metav1.GetOptions{})
		printError(err, cmd, "Error:")
		cmd.Println("This is the secret ", secret)
		cmd.Println("===============================================")
		cmd.Printf("meta: %+v\n\n", objMeta)
		cmd.Printf("Type: %+v\n\n", objTypeMeta)
		cmd.Println("function for secret:", createOutputSecret(secret))
		saveToFile(createOutputSecret(secret), out)
		if fl != "" && out != "" && ns != "" {
			cmd.Printf("Saving %s secret to: %s in %s namespace", convertToBase64(fl), out, ns)
		} else {
			cmd.Println("No file location chosen")
			os.Exit(1)
		}
	},
}
