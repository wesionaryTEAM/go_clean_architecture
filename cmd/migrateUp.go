package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate up cmd",
	Long:  `Command to install version 1 of our application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate up command")
	},
}

func init(){
	migrateCmd.AddCommand(migrateUpCmd)
}
