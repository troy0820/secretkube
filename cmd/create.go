package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
)

func convertToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Creates secret for cluster",
		Long:  `Creates a kubernetes secret for your cluster`,
		Run: func(cmd *cobra.Command, args []string) {
			str := "This is a string"
			str2 := convertToBase64(str)
			fmt.Println("This is a string base64 encoded", str2)
		},
	}
}
