package cmd

import (
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
	"github.com/spf13/cobra"
)
type MigrateCommands struct {
	migration infrastructure.Migrations
	logger lib.Logger
}
func NewMigrateCommands(
	migration infrastructure.Migrations,
	logger lib.Logger,
) MigrateCommands {
	return MigrateCommands{
		migration: migration,
		logger:logger,

	}
}
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate cmd is used for database migration",
	Long:  `migrate cmd is used for database migration: migrate < up >`,
}

func (mc MigrateCommands) Migrate() {
	mc.logger.Info("------- 🤖 Running migration via cobra 🤖 (CLI) -------")
	mc.MigrateUp()
	rootCmd.AddCommand(migrateCmd)
}

