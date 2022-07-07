package cmd

import (
	"log"

	"github.com/edanko/moses/internal/profiles"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen (launch)",
	Short: "Import profiles from gen to db",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("No launch specified")
		}
		profiles.ImportGen(args[0])
	},
}

func init() {
	profilesCmd.AddCommand(genCmd)
}
