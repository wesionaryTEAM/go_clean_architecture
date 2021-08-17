package main

import (
	"clean-architecture/bootstrap"
	"clean-architecture/lib"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()

	logger := lib.GetLogger()
	fx.New(bootstrap.Module, fx.Logger(logger.GetFxLogger())).Run()

}
