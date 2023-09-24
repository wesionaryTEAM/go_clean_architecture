package bootstrap

import (
	"clean-architecture/domain"
	"clean-architecture/pkg"
	"clean-architecture/seeds"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	pkg.Module,
	domain.Module,
	seeds.Module,
)
