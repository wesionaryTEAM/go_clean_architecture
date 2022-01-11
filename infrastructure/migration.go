package infrastructure

import (
	"clean-architecture/lib"
	"fmt"

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

// Migrate -> migrates all table
func (m Migrations) Migrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migration/",
	}

	sqlDB, err := m.db.DB.DB()
	if err != nil {
		m.logger.Error("error in migration", err.Error())
		m.logger.Panic(err)
	}

	_, err = migrate.Exec(sqlDB, "mysql", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error in migration: ", err)
	}
}
