package commands

import (
	"clean-architecture/lib"

	"github.com/spf13/cobra"
)

type TestDBSetupCommand struct{}

// Short returns string about short description of the command
// the string is shown in help screen of cobra command
func (t *TestDBSetupCommand) Short() string {
	return "setup test db command"
}

// Setup is used to setup flags or pre-run steps for the command.
//
// For example,
//  cmd.Flags().IntVarP(&r.num, "num", "n", 5, "description")
//
func (t *TestDBSetupCommand) Setup(cmd *cobra.Command) {

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
func (t *TestDBSetupCommand) Run() lib.CommandRunner {
	return func(l lib.Logger) {
		l.Info("test db setup command")
	}
}

// NewTestDBSetupCommand setup test db command
func NewTestDBSetupCommand() *TestDBSetupCommand {
	return &TestDBSetupCommand{}
}
