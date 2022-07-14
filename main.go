package main

import (
	"clean-architecture/bootstrap"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	_ = bootstrap.RootApp.Execute()
}
