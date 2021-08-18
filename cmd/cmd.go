package cmd

import (
	"clean-architecture/cmd/cli"

	"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRootCommand),
	cli.Module,
)
