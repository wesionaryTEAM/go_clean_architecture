package infrastructure

import (
	"clean-architecture/lib"

	migrate "github.com/rubenv/sql-migrate"
)

// Migrations -> Migration Struct
type Migrations struct {
	logger lib.Logger
	db     Database
}

// NewMigrations -> return new Migrations struct
func NewMigrations(
	logger lib.Logger,
	db Database,
) Migrations {
	return Migrations{
		logger: logger,
		db:     db,
	}
}

// Migrate migrates all migrations that are defined
func (m Migrations) Migrate() error {

	migrations := &migrate.FileMigrationSource{
		Dir: "migration/",
	}

	sqlDB, err := m.db.DB.DB()
	if err != nil {
		return err
	}

	m.logger.Info("running migration.")
	_, err = migrate.Exec(sqlDB, "mysql", migrations, migrate.Up)
	if err != nil {
		return err
	}
	m.logger.Info("migration completed.")
	return nil
}

// RunMigration runs the migration provided logger and database instance
func RunMigration(logger lib.Logger, db Database) error {
	m := Migrations{logger, db}
	return m.Migrate()
}
