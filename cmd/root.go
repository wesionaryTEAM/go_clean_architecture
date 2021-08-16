package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
type RootCommands struct {
	migrateCommands MigrateCommands
}

func NewRootCommands(migrateCommands MigrateCommands) RootCommands {
	return RootCommands{
		migrateCommands: migrateCommands,
	}
}

var rootCmd = &cobra.Command{
	Use:   "clean-architecture",
	Short: "Root command for our application",
	Long:  `Root command for our application, the main purpose is to help setup subcommands`,

}

func (rc RootCommands) Execute() {
	rc.migrateCommands.Migrate()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}




