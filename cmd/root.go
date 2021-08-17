package cmd

import (
	"clean-architecture/lib"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)
type RootCommands struct {
	migrateCommands MigrateCommands
	logger lib.Logger
}

func NewRootCommands(migrateCommands MigrateCommands, logger lib.Logger) RootCommands {
	return RootCommands{
		migrateCommands: migrateCommands,
		logger: logger,
	}
}

var rootCmd = &cobra.Command{
	Use:   "clean-architecture",
	Short: "Root command for our application",
	Long:  `Root command for our application`,

}

func (rc RootCommands) Execute() {
	rc.logger.Info("Running root command of our application")
	rc.migrateCommands.Migrate()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}




