package infrastructure

import (
	"clean-architecture/lib"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//Migrations -> Migration Struct
type Migrations struct {
	logger lib.Logger
	db     Database
}

//NewMigrations -> return new Migrations struct
func NewMigrations(
	logger lib.Logger,
	db Database,
) Migrations {
	return Migrations{
		logger: logger,
		db:     db,
	}
}

//Migrate -> migrates all table
func (m Migrations) Migrate() {
	migrations, err := migrate.New("file://migration/", "mysql://"+m.db.dsn)
	if err != nil {
		m.logger.Error("error in migration", err.Error())
		m.logger.Panic(err)
	}

	m.logger.Info("--- Running Migration ---")
	err = migrations.Steps(1000)
	if err != nil {
		fmt.Println("Error in migration: ", err)
	}
}
