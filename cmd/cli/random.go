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

func NewRandomCommand(logger lib.Logger) RandomCommand {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Prints random string of characters",
	}
	random := RandomCommand{Command: cmd, logger: logger}
	return random
}

func (t *RandomCommand) GetCommand() *cobra.Command {
	return t.Command
}

func (t *RandomCommand) Init() {
	t.Command.Flags().IntVarP(&t.num, "num", "n", 5, "number of random characters to print")
}

func (t *RandomCommand) Run(cmd *cobra.Command, args []string) {
	t.logger.Info("running random command")
	rand.Seed(time.Now().Unix())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, t.num)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	t.logger.Info(string(b))
}
