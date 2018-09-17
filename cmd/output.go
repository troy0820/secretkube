package cmd

import (
	"github.com/spf13/cobra"
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
		if fl != "" {
			cmd.Println("Saving to: ", fl)
		} else {
			cmd.Println("No file location chosen")
			os.Exit(1)
		}
	},
}
