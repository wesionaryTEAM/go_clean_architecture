package commands

import (
	"clean-architecture/lib"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

type RandomCommand struct {
	num int
}

func (r *RandomCommand) Short() string {
	return "generate a random command"
}

func (r *RandomCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&r.num, "num", "n", 5, "length of random number to generate")
}

func (r *RandomCommand) Run() lib.CommandRunner {
	return func(l lib.Logger) {
		l.Info("running random command")
		rand.Seed(time.Now().Unix())
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		b := make([]rune, r.num)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))] //nolint:gosec // for better performance
		}
		l.Info(string(b))
	}
}

func NewRandomCommand() *RandomCommand {
	return &RandomCommand{}
}
