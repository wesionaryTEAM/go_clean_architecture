package domain

import (
	"clean-architecture/domain/features"

	"go.uber.org/fx"
)

var Module = fx.Module("domain",
	fx.Options(
		features.Module,
	),
)
