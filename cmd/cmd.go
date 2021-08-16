package cmd



import (
"go.uber.org/fx"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRootCommands),
	fx.Provide(NewMigrateCommands),
)
