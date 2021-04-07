package cli

import (
	"prototype2/lib"

	"github.com/manifoldco/promptui"
)

// Command has a command
type Command interface {
	Run()
	Name() string
}

// Application cli application
type Application struct {
	logger   lib.Logger
	commands []Command
}

// NewApplication creates new cli application
func NewApplication(
	logger lib.Logger,
) Application {
	return Application{
		logger: logger,
	}
}

// Start starts cli application
func (c Application) Start() {
	c.logger.Zap.Info("⛑  Start CLI...")
	names := []string{}
	commandMap := map[string]Command{}

	for _, command := range c.commands {
		names = append(names, command.Name())
		commandMap[command.Name()] = command
	}

	names = append(names, "EXIT_APPLICATION")

	prompt := promptui.Select{
		Label: "Select the command to run",
		Items: names,
	}

	_, result, err := prompt.Run()

	if err != nil {
		c.logger.Zap.Error("prompt failed")
	}

	if result == "EXIT_APPLICATION" {
		c.logger.Zap.Info("CLI Application Exited")
		return
	}

	commandMap[result].Run()

	c.Start()

}
