package infrastructure

import "os"

// Env has environment stored
type Env struct {
	ServerPort  string
	Environment string
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	SentryDSN   string

	StorageBucketName string

	MailClientID     string
	MailClientSecret string
	MailTokenType    string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("SERVER_PORT")
	env.Environment = os.Getenv("ENVIRONMENT")

	env.DBUsername = os.Getenv("DB_USER")
	env.DBPassword = os.Getenv("DB_PASS")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBName = os.Getenv("DB_NAME")

	env.SentryDSN = os.Getenv("SENTRY_DSN")

	env.StorageBucketName = os.Getenv("STORAGE_BUCKET_NAME")

	env.MailClientID = os.Getenv("MAIL_CLIENT_ID")
	env.MailClientSecret = os.Getenv("MAIL_CLIENT_SECRET")
	env.MailTokenType = os.Getenv("MAIL_TOKEN_TYPE")
}
