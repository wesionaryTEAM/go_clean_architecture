package framework

import (
	"github.com/spf13/viper"
)

type AdminConfig struct {
	Email    string `mapstructure:"ADMIN_EMAIL"`
	Password string `mapstructure:"ADMIN_PASSWORD"`
}

type AWSConfig struct {
	Region            string `mapstructure:"AWS_REGION"`
	AccessKey         string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey   string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	StorageBucketName string `mapstructure:"STORAGE_BUCKET_NAME"`
	CognitoClientID   string `mapstructure:"COGNITO_CLIENT_ID"`
	CognitoUserPoolID string `mapstructure:"COGNITO_USER_POOL_ID"`
}

type SentryConfig struct {
	DSN string `mapstructure:"SENTRY_DSN"`
}

type ServerConfig struct {
	Port        string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	TimeZone    string `mapstructure:"TIMEZONE"`
}

type DatabaseConfig struct {
	Username    string `mapstructure:"DB_USER"`
	Password    string `mapstructure:"DB_PASS"`
	Host        string `mapstructure:"DB_HOST"`
	Port        string `mapstructure:"DB_PORT"`
	Name        string `mapstructure:"DB_NAME"`
	Type        string `mapstructure:"DB_TYPE"`
	ForwardPort string `mapstructure:"DB_FORWARD_PORT"`
}

type Env struct {
	Server ServerConfig `mapstructure:",squash"`

	Database DatabaseConfig `mapstructure:",squash"`

	Sentry             SentryConfig `mapstructure:",squash"`
	MaxMultipartMemory int64        `mapstructure:"MAX_MULTIPART_MEMORY"`

	Admin AdminConfig `mapstructure:",squash"`

	AWS AWSConfig `mapstructure:",squash"`
}

var globalEnv = Env{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetEnv() Env {
	return globalEnv
}

func NewEnv(logger Logger) *Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("cannot read cofiguration", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		logger.Fatal("environment cant be loaded: ", err)
	}

	return &globalEnv
}
