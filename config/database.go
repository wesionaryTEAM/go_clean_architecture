package config

import (
  "fmt"
  "log"
	"os"

  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host string
	Port string
	User string
	DBName string
	Password string
	DBType string
}

func buildDBConfig(host, port, user, name, password string, dbType string) *DBConfig {
	dbConfig := DBConfig{
    Host:     host,
    Port:     port,
    User:     user,
    DBName:   name,
    Password: password,
    DBType:   dbType,
  }
  return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
    "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    dbConfig.User,
    dbConfig.Password,
    dbConfig.Host,
    dbConfig.Port,
    dbConfig.DBName,
  )
}

func InitDatabase() (DB *gorm.DB) {
	dbUser := os.Getenv("DB_USER")
  dbPassword := os.Getenv("DB_PASSWORD")
  dbPort := os.Getenv("DB_PORT")
  dbHost := os.Getenv("DB_HOST")
  dbName := os.Getenv("DB_NAME")
	dbType := os.Getenv("DB_TYPE")
	
	dbConfig := buildDBConfig(dbHost, dbPort, dbUser, dbName, dbPassword, dbType)

  dbURL := dbURL(dbConfig)

	db, err := gorm.Open("mysql", dbURL)
	
	if err != nil {
    fmt.Printf("Cannot connect to database. Host name: %s", dbHost)
    log.Fatal("This is the error:", err)
  } else {
    fmt.Printf("We are connected to the %s database", "mysql")
	}
	
	// db.AutoMigrate()
  return db
}