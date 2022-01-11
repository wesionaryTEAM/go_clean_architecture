package cli

import (
	"clean-architecture/lib"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// RandomCommand test command
type RandomCommand struct {
	*cobra.Command
	logger lib.Logger
	num    int
}

// NewRandomCommand creates new random command
func NewRandomCommand(logger lib.Logger) RandomCommand {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Prints random string of characters",
	}
	random := RandomCommand{Command: cmd, logger: logger}
	return random
}

// GetCommand get the command
func (t *RandomCommand) GetCommand() *cobra.Command {
	return t.Command
}

// Init initialize random command
func (t *RandomCommand) Init() {
	t.Command.Flags().IntVarP(&t.num, "num", "n", 5, "number of random characters to print")
}

// Run run the command
func (t *RandomCommand) Run(cmd *cobra.Command, args []string) {
	t.logger.Info("running random command")
	rand.Seed(time.Now().Unix())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, t.num)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))] //nolint:gosec // for faster performance
	}
	t.logger.Info(string(b))
}
