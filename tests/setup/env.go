package setup

import (
	"clean-architecture/lib"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func TestEnvReplacer(l lib.Logger, env *lib.Env) *lib.Env {
	dbUsername := viper.GetString("TEST_DB_USER")
	if dbUsername != "" {
		env.DBUsername = dbUsername
	}

	dbPassword := viper.GetString("TEST_DB_PASS")
	if dbPassword != "" {
		env.DBPassword = dbPassword
	}

	dbHost := viper.GetString("TEST_DB_HOST")
	if dbHost != "" {
		env.DBHost = dbHost
	}

	dbPort := viper.GetString("TEST_DB_PORT")
	if dbPort != "" {
		env.DBPort = dbPort
	}

	dbName := viper.GetString("TEST_DB_NAME")
	fmt.Println(os.Args[0])
	if dbName != "" {
		env.DBName = dbName
	}

	caller := GetCaller()

	env.DBName = strings.ReplaceAll(fmt.Sprintf("TEST_%s_%s", caller, env.DBName), ".", "_")

	dbType := viper.GetString("TEST_DB_TYPE")
	if dbType != "" {
		env.DBType = dbType
	}

	l.Info("Test Environment: %+v", env)
	return env
}
