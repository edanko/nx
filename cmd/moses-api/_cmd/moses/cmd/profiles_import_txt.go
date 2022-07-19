package cmd

import (
	"log"

	"github.com/edanko/moses/internal/profiles"

	"github.com/spf13/cobra"
)

var txtCmd = &cobra.Command{
	Use:   "txt (launch)",
	Short: "Import profiles from txt to db",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("No launch specified")
		}
		profiles.ImportTxt(args[0])
	},
}

func init() {
	profilesCmd.AddCommand(txtCmd)
}
