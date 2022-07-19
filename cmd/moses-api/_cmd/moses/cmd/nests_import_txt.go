package cmd

import (
	"log"

	"github.com/edanko/moses/internal/nests"

	"github.com/spf13/cobra"
)

var nestsImportTxtCmd = &cobra.Command{
	Use:   "txt",
	Short: "Import profiles and nests from txt to db",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("No launch specified")
		}
		nests.ImportTxt(args[0])
	},
}

func init() {
	nestsImportCmd.AddCommand(nestsImportTxtCmd)
}
