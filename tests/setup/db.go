package setup

import (
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
	"log"
	"testing"

	"go.uber.org/fx"
)

func TeardownDB() {
	// because actual instance is not passable as
	// this function is being called from TestMain func
	t := testing.T{}

	var db infrastructure.Database
	var env *lib.Env
	var l lib.Logger

	_, cancel, err := DI(&t,
		fx.Options(
			fx.Populate(&db),
			fx.Populate(&env),
			fx.Populate(&l),
		),
	)
	defer cancel()
	if err != nil {
		log.Println(err)
		return
	}

	err = db.Exec("DROP DATABASE IF EXISTS " + env.DBName).Error
	if err != nil {
		l.Fatalf("couldn't teardown database: %s", err)
	}
	l.Info("test database teardown successful")
}
