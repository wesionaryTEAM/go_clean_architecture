package main

import (
	"clean-architecture/cmd"
)

func main() {
	//godotenv.Load()
	//
	//logger := lib.GetLogger()
	//fx.New(bootstrap.Module, fx.Logger(logger.GetFxLogger())).Run()
	cmd.Execute()
}
