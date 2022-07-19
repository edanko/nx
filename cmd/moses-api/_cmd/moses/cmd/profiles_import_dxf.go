package cmd

import (
	"log"

	"github.com/edanko/moses/internal/profiles"

	"github.com/spf13/cobra"
)

var dxfCmd = &cobra.Command{
	Use:   "dxf (launch)",
	Short: "Import profiles from dxf to db",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("No launch specified")
		}
		profiles.ImportDxf(args[0])
	},
}

func init() {
	profilesCmd.AddCommand(dxfCmd)
}
