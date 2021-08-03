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
	dsn string
}

// NewDatabase creates a new database instance
func NewDatabase(logger lib.Logger, env lib.Env) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)

	if env.Environment != "local" {
		url = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.DBUsername,
			env.DBPassword,
			env.DBHost,
			env.DBName,
		)
	}

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	logger.Info("Database connection established")

	return Database{
		DB:  db,
		dsn: url,
	}
}
