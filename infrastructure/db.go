package infrastructure

import (
	"clean-architecture/lib"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database modal
type Database struct {
	*gorm.DB
	dsn string
}

// NewDatabase creates a new database instance
func NewDatabase(Zaplogger lib.Logger, env lib.Env) Database {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	Zaplogger.Info(env)

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

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
	if err != nil {
		Zaplogger.Info("Url: ", url)
		Zaplogger.Panic(err)
	}

	Zaplogger.Info("Database connection established")

	return Database{
		DB: db,
		dsn: url,
	}
}
