package commands

import (
	"clean-architecture/lib"

	"github.com/spf13/cobra"
)


var VU int
var DU int

type LoadTestCommand struct{

}

func NewLoadTestingCommand() *LoadTestCommand{
	return &LoadTestCommand{}
}

func(ltc *LoadTestCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&VU,"virtual-user","v",1,"A falg that sets total virtual users")
	cmd.Flags().IntVarP(&DU,"duration","d",1,"A flag that sets total duration")

}
func (ltc *LoadTestCommand) Run(cmd *cobra.Command, args []string) lib.CommandRunner {
	return func(l lib.Logger) {
		l.Info("running loadtest command",args)
		vus,_:=cmd.Flags().GetInt("virtual-user")
		dus,_:=cmd.Flags().GetInt("duration")
		l.Info("virtual users",vus, "duration",dus)
	}
}
func (ltc *LoadTestCommand) Short() string{
	return "Shell for loadtesting"
}