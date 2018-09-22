package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes/fake"
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

		if fl != "" && out != "" && ns != "" {
			cmd.Printf("Saving %s secret to: %s in %s namespace", convertToBase64(fl), out, ns)
		} else {
			cmd.Println("No file location chosen")
			os.Exit(1)
		}
	},
}
