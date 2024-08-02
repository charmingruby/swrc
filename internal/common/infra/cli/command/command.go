package command

import "github.com/spf13/cobra"

func NewCommand(rootCommand *cobra.Command) *Command {
	return &Command{
		rootCommand: rootCommand,
	}
}

type Command struct {
	rootCommand *cobra.Command
}

func (cmd *Command) Register() {
	cmd.rootCommand.AddCommand(cmd.newHealthCheckCommand())
}
