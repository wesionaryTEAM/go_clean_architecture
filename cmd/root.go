package cmd

import (
	"clean-architecture/cmd/cli"
	"clean-architecture/lib"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clean-architecture",
	Short: "Commander for clean architecture",
	Long: `
		This is a command runner or cli for api architecture in golang. 
		Using this we can use underlying dependency injection container for running scripts. 
		Main advantage is that, we can use same services, repositories, infrastructure present in the application itself`,
	TraverseChildren: true,
}

// Command command interface
type Command interface {

	// Init initializes the command by default only run commands are initilialized
	Init()

	// GetCommand gets the underlying cobra instance
	GetCommand() *cobra.Command

	// Run runs the command
	Run(cmd *cobra.Command, args []string)
}

// RootCommand root of the application
type RootCommand struct {
	*cobra.Command
	logger   lib.Logger
	commands []Command
}

// NewRootCommand creates new root command
func NewRootCommand(
	logger lib.Logger,
	randomCmd cli.RandomCommand,
) RootCommand {
	cmd := RootCommand{
		Command: rootCmd,
		logger:  logger,
		commands: []Command{
			&randomCmd,
		},
	}
	cmd.InitCommands()
	return cmd
}

// InitCommands initializes the command and sub-commands
func (r RootCommand) InitCommands() {
	for _, c := range r.commands {
		cmd := c.GetCommand()
		if cmd != nil {
			cmd.Run = c.Run
			c.Init()
		}
	}

	for _, c := range r.commands {
		cmd := c.GetCommand()
		if cmd != nil {
			rootCmd.AddCommand(cmd)
		}
	}
}
