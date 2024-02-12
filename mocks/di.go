package mocks

import (
	"clean-architecture/domain"
	"clean-architecture/pkg"
	"clean-architecture/pkg/framework"
	"context"

	"github.com/onsi/ginkgo/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap/zaptest"
)

func DI(t ginkgo.GinkgoTInterface, opts ...fx.Option) error {
	finalOpts := []fx.Option{
		pkg.Module,
		domain.Module,
		fx.Decorate(
			NewMockDB,
			func() framework.Logger {
				return framework.Logger{
					SugaredLogger: zaptest.NewLogger(t).Sugar(),
				}
			},
		),
	}
	finalOpts = append(finalOpts, opts...)

	app := fxtest.New(t, finalOpts...)

	ctx := context.Background()
	return app.Start(ctx)
}
