package infrastructure

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//Migrations -> Migration Struct
type Migrations struct {
	logger Logger
	env    Env
}

//NewMigrations -> return new Migrations struct
func NewMigrations(
	logger Logger,
	env Env,
) Migrations {
	return Migrations{
		logger: logger,
		env:    env,
	}
}

//Migrate -> migrates all table
func (m Migrations) Migrate() {
	m.logger.Info("Migrating schemas...")

	USER := m.env.DBUsername
	PASS := m.env.DBPassword
	HOST := m.env.DBHost
	PORT := m.env.DBPort
	DBNAME := m.env.DBName
	ENVIRONMENT := m.env.Environment

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	if ENVIRONMENT != "local" {
		dsn = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			USER,
			PASS,
			HOST,
			DBNAME,
		)
	}

	migrations, err := migrate.New("file://migration/", "mysql://"+dsn)
	if err != nil {
		m.logger.Error("error in migration", err.Error())
	}

	m.logger.Info("--- Running Migration ---")
	err = migrations.Steps(1000)
	if err != nil {
		m.logger.Error("Error in migration: ", err.Error())
	}
}
