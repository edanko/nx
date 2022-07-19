package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nestsExportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
	},
}

func init() {
	nestsCmd.AddCommand(nestsExportCmd)
}
