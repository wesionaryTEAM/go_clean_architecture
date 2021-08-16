package cmd

import (
	"clean-architecture/infrastructure"
	"github.com/spf13/cobra"
)
type MigrateCommands struct {

	migration infrastructure.Migrations
}
func NewMigrateCommands(

	migration infrastructure.Migrations,
) MigrateCommands {
	return MigrateCommands{
		migration: migration,

	}
}
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate cmd is used for database migration",
	Long:  `migrate cmd is used for database migration: migrate < up | down >`,
}

func (mc MigrateCommands) Migrate() {
	mc.MigrateUp()
	rootCmd.AddCommand(migrateCmd)
}

