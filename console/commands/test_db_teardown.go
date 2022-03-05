package commands

import (
	"clean-architecture/infrastructure"
	"clean-architecture/lib"

	"github.com/spf13/cobra"
)

type TestDBTeardownCommand struct{}

func (t *TestDBTeardownCommand) Short() string {
	return "test database teardown (drop database)"
}

func (t *TestDBTeardownCommand) Setup(cmd *cobra.Command) {}

func (t *TestDBTeardownCommand) PreRun(cmd *cobra.Command, args []string) {
	lib.ForceTestOverride()
}

func (t *TestDBTeardownCommand) Run() lib.CommandRunner {
	return func(l lib.Logger, db infrastructure.Database, env lib.DBEnv) {
		l.Info("database: ", env.Name)
		err := db.Exec("DROP DATABASE IF EXISTS " + env.Name).Error
		if err != nil {
			l.Fatalf("couldn't teardown database: %s", err)
		}
		l.Info("test database teardown successfull")
	}
}

// NewTestDBTeardownCommand create test db teardown
func NewTestDBTeardownCommand() *TestDBTeardownCommand {
	return &TestDBTeardownCommand{}
}
