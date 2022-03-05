package lib

import (
	"clean-architecture/utils"

	"github.com/spf13/viper"
)

var _testOverride = false

// ForceTestOverride forces test environment
// when overridden test database will be used
func ForceTestOverride() {
	_testOverride = true
}

// DBEnv has information related to database environment variables
type DBEnv struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	Type     string
}

type Env struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBType     string `mapstructure:"DB_TYPE"`

	TestDBUsername string `mapstructure:"TEST_DB_USER"`
	TestDBPassword string `mapstructure:"TEST_DB_PASS"`
	TestDBHost     string `mapstructure:"TEST_DB_HOST"`
	TestDBPort     string `mapstructure:"TEST_DB_PORT"`
	TestDBName     string `mapstructure:"TEST_DB_NAME"`
	TestDBType     string `mapstructure:"TEST_DB_TYPE"`

	MailClientID     string `mapstructure:"MAIL_CLIENT_ID"`
	MailClientSecret string `mapstructure:"MAIL_CLIENT_SECRET"`
	MailTokenType    string `mapstructure:"MAIL_TOKEN_TYPE"`

	SentryDSN          string `mapstructure:"SENTRY_DSN"`
	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
	StorageBucketName  string `mapstructure:"STORAGE_BUCKET_NAME"`
}

var globalEnv = Env{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetEnv() Env {
	return globalEnv
}

func NewEnv(logger Logger) *Env {

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("cannot read cofiguration", err)
	}

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		logger.Fatal("environment cant be loaded: ", err)
	}

	return &globalEnv
}

func NewDBEnv(env *Env) DBEnv {
	if utils.IsTestEnv() || _testOverride {
		return DBEnv{
			Username: env.TestDBUsername,
			Password: env.TestDBPassword,
			Host:     env.TestDBHost,
			Port:     env.TestDBPort,
			Name:     env.TestDBName,
			Type:     env.TestDBType,
		}
	}
	return DBEnv{
		Username: env.DBUsername,
		Password: env.DBPassword,
		Host:     env.DBHost,
		Port:     env.DBPort,
		Name:     env.DBName,
		Type:     env.DBType,
	}
}
