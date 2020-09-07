package main

import (
	"prototype2/infrastructure"
	"prototype2/utils"
)

func main() {
	utils.LoadEnv()

	err := utils.SetupLumberjackLoging()
	if err != nil {
		panic(err.Error())
	}


	utils.SetupSentry()

	db := infrastructure.SetupModels()

	fb := infrastructure.InitializeFirebase()

	infrastructure.SetupRoutes(db, fb)
}
