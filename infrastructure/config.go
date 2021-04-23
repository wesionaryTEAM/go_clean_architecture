package infrastructure

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	New,
)

type Spec struct {
	Env         string `envconfig:"ENV" default:"production" required:"true"`
	HTTPPort    string `envconfig:"PORT" default:"8000" required:"true"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"INFO" required:"true"`
	LogEncType  string `envconfig:"LOG_ENC_TYPE" default:"pretty" required:"true"`
	PostgresDSN string `envconfig:"DATABASE_URL" required:"true"`

	SentryDSN string `envconfig:"SENTRY_DSN"`
	SentryENV string `envconfig:"SENTRY_ENVIRONMENT"`
}

func New() (*Spec, error) {
	var cfg Spec
	if err := envconfig.Process("SKELETON", &cfg); err != nil {
		return nil, fmt.Errorf("cannot init config: %w", err)
	}

	return &cfg, nil
}
