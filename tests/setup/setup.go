package setup

import (
	"clean-architecture/lib"
	"context"
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func SetupDI(t *testing.T, option fx.Option) (context.Context, context.CancelFunc, error) {
	app := fxtest.New(t,
		fx.Options(
			fx.WithLogger(func(l lib.Logger) fxevent.Logger {
				return l.GetFxLogger()
			}),
			TestModule,
			option,
		),
	)
	ctx, cancel := context.WithCancel(context.Background())
	err := app.Start(ctx)
	return ctx, cancel, err
}
