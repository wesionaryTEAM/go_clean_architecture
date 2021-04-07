package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"prototype2/bootstrap"
)

func main() {
	godotenv.Load()

	fx.New(bootstrap.Module).Run()
}
