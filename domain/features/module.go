package features

import (
	"clean-architecture/domain/features/user"

	"go.uber.org/fx"
)

var Module = fx.Module("features",
	fx.Options(
		user.Module,
	),
)
