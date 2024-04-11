package commands

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/seeds"
	"errors"

	"github.com/spf13/cobra"
)

type SeedCommand struct {
	names  []string
	runAll bool
}

func (s *SeedCommand) Short() string {
	return "run seed command"
}

func NewSeedCommand() *SeedCommand {
	return &SeedCommand{}
}

func (s *SeedCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().StringArrayVarP(
		&s.names,
		"name",
		"n",
		[]string{},
		"name of the seed to run (can be used multiple times)",
	)
	cmd.Flags().BoolVar(&s.runAll, "all", false, "run all seeds")
}

func (s *SeedCommand) Run() framework.CommandRunner {
	return func(
		l framework.Logger,
		seeds seeds.Seeds,
	) {

		// run all seeds
		if s.runAll {
			for _, seed := range seeds {
				seed.Setup()
			}
			return
		}

		// validate names array
		if len(s.names) == 0 {
			l.Info("no seed name provided")
			return
		}

		// run selected seeds
		for _, name := range s.names {
			if err := runSeed(name, &seeds); err != nil {
				l.Infof("Error running %s: %s", name, err)
			}
		} // end for loop
	} // end return func
} // end run func

func runSeed(name string, seeds *seeds.Seeds) error {
	isValid := false
	for _, seed := range *seeds {
		if name == seed.Name() {
			isValid = true
			seed.Setup()
		}
	} // end for loop

	if !isValid {
		return errors.New("invalid seed name")
	}
	return nil
} // end runSeed func
