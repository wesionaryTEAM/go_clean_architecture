package infrastructure

import (
	"clean-architecture/pkg/framework"
	"fmt"
	"regexp"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
	logger framework.Logger
	env    *framework.Env
}

// validDBIdentifierRegex matches valid database identifiers (must start with letter or underscore,
// followed by alphanumeric characters or underscores)
var validDBIdentifierRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// maxDBNameLength defines the maximum allowed database name length (PostgreSQL has the strictest limit at 63)
const maxDBNameLength = 63

func validateDBName(dbName string) error {
	if dbName == "" {
		return fmt.Errorf("database name cannot be empty")
	}
	if len(dbName) > maxDBNameLength {
		return fmt.Errorf("database name exceeds maximum length of %d characters", maxDBNameLength)
	}
	if !validDBIdentifierRegex.MatchString(dbName) {
		return fmt.Errorf("database name contains invalid characters: must start with a letter or underscore, followed by alphanumeric characters or underscores")
	}
	return nil
}

func NewDatabase(logger framework.Logger, env *framework.Env) (*Database, error) {
	if err := validateDBName(env.Database.Name); err != nil {
		return nil, fmt.Errorf("invalid database name: %w", err)
	}
	// Determine SSL mode based on the environment
	sslMode := "require"
	if env.Server.Environment == "local" || env.Server.Environment == "workflow" {
		sslMode = "disable"
	}
	url := fmt.Sprintf("host=%s port=%s user=%s password=%s database=postgres sslmode=%s", env.Database.Host, env.Database.Port, env.Database.Username, env.Database.Password, sslMode)
	logger.Info("connecting to database (postgres)")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.GetGormLogger(), TranslateError: true})
	if err != nil {
		return nil, err
	}
	logger.Info("checking if the target database exists (postgres)")
	var exists bool
	if err = db.Raw("SELECT 1 FROM pg_database WHERE datname = ?", env.Database.Name).Scan(&exists).Error; err != nil {
		return nil, fmt.Errorf("failed to check if the database exists: %w", err)
	}
	if !exists {
		logger.Info("creating target database (postgres)")
		if err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", env.Database.Name)).Error; err != nil {
			return nil, fmt.Errorf("couldn't create database: %w", err)
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Info("couldn't get db connection (postgres)")
		return nil, fmt.Errorf("failed to get db connection for closing: %w", err)
	}
	if dbErr := sqlDB.Close(); dbErr != nil {
		logger.Info("couldn't close db connection (postgres)")
		return nil, fmt.Errorf("failed to close db connection: %w", dbErr)
	}

	// connect to the target database
	url = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", env.Database.Host, env.Database.Port, env.Database.Username, env.Database.Password, env.Database.Name, sslMode)
	logger.Info("connecting to target database (postgres)")
	db, err = gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.GetGormLogger(), TranslateError: true})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to target database: %w", err)
	}
	logger.Info("database connection established (postgres)")
	database := Database{
		DB:     db,
		logger: logger,
		env:    env,
	}
	return &database, nil
}
