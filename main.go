package main

import (
	"clean-architecture/bootstrap"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()

	fx.New(bootstrap.Module).Run()
}
