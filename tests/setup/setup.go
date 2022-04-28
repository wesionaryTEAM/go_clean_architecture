package setup

import (
	"clean-architecture/api/middlewares"
	"clean-architecture/api/routes"
	"clean-architecture/lib"
	"context"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

var _caller string

// GetCaller get's the test caller
func GetCaller() string {
	return _caller
}

func DI(t fxtest.TB, option fx.Option) (context.Context, context.CancelFunc, error) {
	var middleware middlewares.Middlewares
	var route routes.Routes
	app := fxtest.New(t,
		fx.Options(
			TestModule,
			fx.WithLogger(func(l lib.Logger) fxevent.Logger {
				return l.GetFxLogger()
			}),
			option,
			fx.Populate(&middleware),
			fx.Populate(&route),
		),
		fx.NopLogger,
	)
	ctx, cancel := context.WithCancel(context.Background())
	err := app.Start(ctx)

	middleware.Setup()
	route.Setup()
	return ctx, cancel, err
}

// init sets up working directory for tests
func init() {
	_, filename, _, _ := runtime.Caller(0)
	_caller = filepath.Base(filename)
	dir := path.Join(path.Dir(filename), "..", "..")
	if err := os.Chdir(dir); err != nil {
		log.Fatal(err)
	}
}
