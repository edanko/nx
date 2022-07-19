package cmd

import (
	"github.com/spf13/cobra"
)

var remnantsCmd = &cobra.Command{
	Use:   "remnants",
	Short: "Subset of remnants-related commands",
	/* Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remnant called")
	}, */
}

func init() {
	rootCmd.AddCommand(remnantsCmd)
}
