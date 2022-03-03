package commands

import (
	"clean-architecture/lib"

	"github.com/spf13/cobra"
)

type TestDBTeardownCommand struct{}

// Short returns string about short description of the command
// the string is shown in help screen of cobra command
func (t *TestDBTeardownCommand) Short() string {
	return "test database teardown (drop database)"
}

// Setup is used to setup flags or pre-run steps for the command.
//
// For example,
//  cmd.Flags().IntVarP(&r.num, "num", "n", 5, "description")
//
func (t *TestDBTeardownCommand) Setup(cmd *cobra.Command) {

}

// Run runs the command runner
// run returns command runner which is a function with dependency
// injected arguments.
//
// For example,
//  Command{
//   Run: func(l lib.Logger) {
// 	   l.Info("i am working")
// 	 },
//  }
//
func (t *TestDBTeardownCommand) Run() lib.CommandRunner {
	return func(l lib.Logger) {
		l.Info("test db teardown command")
	}
}

// NewTestDBTeardownCommand create test db teardown
func NewTestDBTeardownCommand() *TestDBTeardownCommand {
	return &TestDBTeardownCommand{}
}
