package infrastructure

import (
	"clean-architecture/pkg/framework"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

// MySQLConnect implements provided MySQL logic.
func MySQLConnect(logger framework.Logger, env *framework.Env) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort)
	logger.Info("opening db connection (mysql)")
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		return nil, err
	}
	logger.Info("creating database if it doesn't exist (mysql)")
	if err = db.Exec("CREATE DATABASE IF NOT EXISTS " + env.DBName).Error; err != nil {
		logger.Info("couldn't create database (mysql)")
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	if dbErr := sqlDb.Close(); dbErr != nil {
		return nil, dbErr
	}
	logger.Info("using given database (mysql)")
	urlWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	db, err = gorm.Open(mysql.Open(urlWithDB), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		return nil, err
	}
	conn, err := db.DB()
	if err != nil {
		logger.Info("couldn't get db connection (mysql)")
		return nil, err
	}
	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(1)
	return db, nil
}

// PostgresConnect implements provided PostgreSQL logic.
func PostgresConnect(logger framework.Logger, env *framework.Env) (*gorm.DB, error) {
	// Determine SSL mode based on the environment
	sslMode := "require"
	if env.Environment == "local" || env.Environment == "workflow" {
		sslMode = "disable"
	}
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s database=postgres sslmode=%s", env.DBHost, env.DBPort, env.DBUsername, env.DBPassword, sslMode)
	logger.Info("Connection to database (postgres)")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.GetGormLogger()})
	if err != nil {
		logger.Info("Url: ", url)
		return nil, err
	}
	logger.Info("checking if the target database exists (postgres)")
	var exists bool
	if err = db.Raw("SELECT 1 FROM pg_database WHERE datname = ?", env.DBName).Scan(&exists).Error; err != nil {
		return nil, fmt.Errorf("failed to check if the database exists: %w", err)
	}
	if !exists {
		logger.Info("creating target database (postgres)")
		if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", env.DBName)).Error; err != nil {
			return nil, fmt.Errorf("couldn't create database: %w", err)
		}
	}
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()

	// connect to the target database
	url = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", env.DBHost, env.DBPort, env.DBUsername, env.DBPassword, env.DBName, sslMode)
	logger.Info("connecting to target database (postgres)")
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.GetGormLogger(), TranslateError: true})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to target database: %w", err)
	}
	logger.Info("database connection established (postgres)")
	return db, nil
}
