package cmd

import (
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/kubernetes/pkg/apis/core"
	"os"
)

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Creates output of the secret",
	Long: `Creates the output of the secret you want to
create.  This output can be saved to a file or printed to the screen`,
	Run: func(cmd *cobra.Command, args []string) {
		fl, err := cmd.Flags().GetString("file")
		if err != nil {
			cmd.Println(err.Error())
		}
		//		f, err := os.Open(fl)
		//		if err != nil {
		//			cmd.Println("File doesn't exist")
		//			os.Exit(1)
		//		}
		out, err := cmd.Flags().GetString("output")
		if err != nil {
			cmd.Println(err.Error())
		}
		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			cmd.Println(err.Error())
		}
		clientset := fake.NewSimpleClientset()
		cmd.Println("clientset", clientset)
		objMeta := metav1.ObjectMeta{
			Name:      "fake-secret",
			Namespace: "default",
		}
		objTypeMeta := metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		}
		sec := core.Secret{}
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
