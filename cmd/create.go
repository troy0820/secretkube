package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
)

func convertToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates secret for cluster",
	Long:  `Creates a kubernetes secret for your cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fl, err := cmd.Flags().GetString("config")
		if err != nil {
			fmt.Println(err.Error())
		}
		var str2 string
		if fl != "" {
			str2 = convertToBase64(fl)
		} else {
			str2 = convertToBase64("This is the default string")
		}
		fmt.Println("This is a string base64 encoded", str2)
	},
}
