package utils

import "go.uber.org/fx"

func FxReplaceAs(impl, iface interface{}) fx.Option {
	return fx.Replace(
		fx.Annotate(impl, fx.As(iface)),
	)
}
