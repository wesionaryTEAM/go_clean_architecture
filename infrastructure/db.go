package infrastructure

import (
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
}

// NewDatabase creates a new database instance
func NewDatabase(Zaplogger Logger, config Env) Database {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	Zaplogger.Zap.Info(config)

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	if config.Environment != "local" {
		url = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DBUsername,
			config.DBPassword,
			config.DBHost,
			config.DBName,
		)
	}

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBName + ";")
	if err != nil {
		Zaplogger.Zap.Info("Url: ", url)
		Zaplogger.Zap.Panic(err)
	}

	Zaplogger.Zap.Info("Database connection established")

	return Database{DB: db}
}
