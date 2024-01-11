package domain

import (
	"clean-architecture/domain/features/user"

	"go.uber.org/fx"
)

var Module = fx.Module("domain",
	fx.Options(
		user.Module,
	),
)
