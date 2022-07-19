package cmd

import (
	"github.com/spf13/cobra"
)

var profilesImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import profiles to db",
}

func init() {
	profilesCmd.AddCommand(profilesImportCmd)
}
