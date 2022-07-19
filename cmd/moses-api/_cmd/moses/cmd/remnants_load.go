package cmd

import (
	"fmt"
	"log"

	"github.com/edanko/moses/internal/remnant"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load remnants from old csv",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("csv file path not specified!")
		}
		total, err := remnant.Load(args[0])
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("remnant records processed: ", total)
	},
}

func init() {
	remnantsCmd.AddCommand(loadCmd)
}
