package infrastructure

import (
	"clean-architecture/lib"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(logger lib.Logger, env *lib.Env) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort)
	if env.DBType != "mysql" {
		url = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/?charset=utf8mb4&parseTime=True&loc=Local",
			env.DBUsername,
			env.DBPassword,
			env.DBHost,
		)
	}

	logger.Info("opening db connection")
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	logger.Info("creating database if it does't exist")
	if err = db.Exec("CREATE DATABASE IF NOT EXISTS " + env.DBName).Error; err != nil {
		logger.Info("couldn't create database")
		logger.Panic(err)
	}

	logger.Info("using given database")
	if err := db.Exec(fmt.Sprintf("USE %s", env.DBName)).Error; err != nil {
		logger.Info("cannot use the given database")
		logger.Panic(err)
	}
	logger.Info("database connection established")

	database := Database{DB: db}
	logger.Info("currentDatabase:", db.Migrator().CurrentDatabase())

	if err := RunMigration(logger, database); err != nil {
		logger.Info("migration failed.")
		logger.Panic(err)
	}

	return Database{DB: db}
}
