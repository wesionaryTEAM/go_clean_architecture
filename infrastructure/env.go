package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort        string `mapstructure:"SERVER_PORT"`
	Environment       string `mapstructure:"ENVIRONMENT"`
	DBUsername        string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBName            string `mapstructure:"DB_NAME"`
	SentryDSN         string `mapstructure:"SENTRY_DSN"`
	StorageBucketName string `mapstructure:"STORAGE_BUCKET_NAME"`
	MailClientID      string `mapstructure:"MAIL_CLIENT_ID"`
	MailClientSecret  string `mapstructure:"MAIL_CLIENT_SECRET"`
	MailTokenType     string `mapstructure:"MAIL_TOKEN_TYPE"`
}

func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read cofiguration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	log.Printf("%#v \n", env)
	return env
}
