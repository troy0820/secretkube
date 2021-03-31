package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func printError(err error, cmd *cobra.Command, msg string) {
	red := color.New(color.FgRed).SprintFunc()
	if err != nil {
		cmd.Println(red(msg), red(err.Error()))
	}
}

func printErrorWithExit(err error, cmd *cobra.Command, msg string) {
	red := color.New(color.FgRed).SprintFunc()
	if err != nil {
		cmd.Println(red(msg), red(err.Error()))
		os.Exit(1)
	}
}

func writeToStdOut(w io.Writer, sec *v1.Secret) error {
	color.Set(color.FgGreen)
	fmt.Fprint(w, createOutputSecret(sec))
	color.Unset()
	return nil
}
func createOutputSecret(sec *v1.Secret) string {
	var a string
	for k, v := range sec.StringData {
		a += fmt.Sprintf("  %s: %s\n", string(k), convertToBase64(string(v)))
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

func createSecret(name string, stringdata map[string]string, data map[string][]byte) (*v1.Secret, error) {
	return &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		Data:       data,
		StringData: stringdata,
		Type:       "Opaque",
	}, nil
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

		m, err := MakeMapFromJSON(fl)
		printErrorWithExit(err, cmd, "Error:")
		var out string
		if cmd.Flags().Changed("output") {
			var err error
			out, err = cmd.Flags().GetString("output")
			printError(err, cmd, "Error:")
		}

		ns := "default"
		if cmd.Flags().Changed("namespace") {
			ns, err = cmd.Flags().GetString("namespace")
			printError(err, cmd, "Error:")
		}

		clientset := fake.NewSimpleClientset()
		bytemap := TurnMapToBytes(m)
		convertMapValuesToBase64(bytemap)
		outputSecret, err := createSecret(name, m, bytemap)
		printError(err, cmd, "Error:")
		returnedSecret, err := clientset.CoreV1().Secrets(ns).Create(outputSecret)
		printError(err, cmd, "Secret:")
		secret, err := clientset.CoreV1().Secrets(ns).Get(returnedSecret.GetName(), metav1.GetOptions{})
		printError(err, cmd, "Error:")
		if out != "" {
			saveToFile(createOutputSecret(secret), out)
			cmd.Printf("Secret saved to %s file \n", out)
			color.Set(color.FgGreen)
			cmd.Println("\nSecret: \n", createOutputSecret(secret))
			color.Unset()
		} else {
			if err := writeToStdOut(os.Stdout, secret); err != nil {
				printError(err, cmd, "Error:")
			}
		}
	},
}
