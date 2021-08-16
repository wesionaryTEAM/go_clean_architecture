package cmd

import (
	"github.com/spf13/cobra"
)


func(cm MigrateCommands) MigrateUp(){
	 migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "migrate up cmd",
		Long:  `Command to install version 1 of our application`,
		Run: func(cmd *cobra.Command, args []string) {
			cm.logger.Info("Running migrate up command")
			cm.migration.Migrate()
		},

	}
	migrateCmd.AddCommand(migrateUpCmd)
}
