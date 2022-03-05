package lib

import (
	"github.com/spf13/viper"
)

type Env struct {
	ServerPort         string `mapstructure:"SERVER_PORT"`
	Environment        string `mapstructure:"ENVIRONMENT"`
	LogLevel           string `mapstructure:"LOG_LEVEL"`
	DBUsername         string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASS"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	TestDBPort         string `mapstructure:"TEST_DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	DBType             string `mapstructure:"DB_TYPE"`
	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
	SentryDSN          string `mapstructure:"SENTRY_DSN"`
	StorageBucketName  string `mapstructure:"STORAGE_BUCKET_NAME"`
	MailClientID       string `mapstructure:"MAIL_CLIENT_ID"`
	MailClientSecret   string `mapstructure:"MAIL_CLIENT_SECRET"`
	MailTokenType      string `mapstructure:"MAIL_TOKEN_TYPE"`
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
