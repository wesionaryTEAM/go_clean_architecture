package infrastructure

import "github.com/spf13/viper"

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

func LoadConfig(path string) (config Env, err error) {

	viper.AddConfigPath(path)

	// to look specfic file name with config name : app
	viper.SetConfigName("app")

	// set config Type that may be env , json, xml and soon other
	viper.SetConfigType("env")

	// to read env values from environment variables, so we can call Automatic.env() to automatically overide values that
	// it has read from the config file with values of coressponding environment variables if they
	// exist in default
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// to unmarshal values into target config object
	err = viper.Unmarshal(&config)
	return
	// loading function is completed
}
