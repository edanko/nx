package cmd

import (
	"log"

	"github.com/edanko/moses/internal/profiles"

	"github.com/spf13/cobra"
)

var csvCmd = &cobra.Command{
	Use:   "csv (launch)",
	Short: "Import profiles from csv to db",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("No launch specified")
		}
		profiles.ImportCsv(args[0])
	},
}

func init() {
	profilesCmd.AddCommand(csvCmd)
}
