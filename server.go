package main

import (
	"prototype2/infrastructure"
	"prototype2/utils"
)

func main() {
	utils.LoadEnv()

	db := infrastructure.SetupModels()

	infrastructure.SetupRoutes(db)
}
