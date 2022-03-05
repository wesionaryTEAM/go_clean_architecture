package infrastructure

import (
	"clean-architecture/lib"
	"clean-architecture/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database modal
type Database struct {
	*gorm.DB
}

func getDSN(env *lib.Env, isTest bool) string {

	dbName := env.DBName
	port := env.DBPort

	if isTest {
		dbName = fmt.Sprintf("%s_test", env.DBName)
		port = env.TestDBPort
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, port, dbName)
	if env.DBType != "mysql" {
		url = fmt.Sprintf(
			"%s:%s@unix(/cloudsql/%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.DBUsername,
			env.DBPassword,
			env.DBHost,
			env.DBName,
		)
	}
	fmt.Println(os.Args[0])
	fmt.Println(url)
	return url
}

// NewDatabase creates a new database instance
func NewDatabase(logger lib.Logger, env *lib.Env) Database {

	url := getDSN(env, false)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	logger.Info("Database connection established")

	database := Database{DB: db}

	if utils.IsTestEnv() {
		database = setupTestDB(db, env, logger)
	}

	// run migration
	if err := RunMigration(logger, database); err != nil {
		logger.Info("migration failed.")
		logger.Panic(err)
	}

	return Database{DB: db}
}

func setupTestDB(db *gorm.DB, env *lib.Env, logger lib.Logger) Database {
	err := db.Exec("CREATE DATABASE IF NOT EXISTS " + fmt.Sprintf("%s_test", env.DBName)).Error
	if err != nil {
		logger.Info("couldn't create test database")
		logger.Panic(err)
	}

	// connect to test database closing the current connection
	sqlDB, err := db.DB()
	if err != nil {
		logger.Info("couldn't get database instance")
		logger.Panic(err)
	}

	if err := sqlDB.Close(); err != nil {
		logger.Info("couldn't close database")
		logger.Panic(err)
	}

	// create connection to test database
	url := getDSN(env, true)
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url: ", url)
		logger.Panic(err)
	}

	return Database{DB: db}
}
