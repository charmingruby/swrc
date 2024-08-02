package cli

import (
	"github.com/charmingruby/swrc/internal/common/infra/cli/command"
	"github.com/spf13/cobra"
)

func NewCLI(rootCommand *cobra.Command) *CLI {
	command := command.NewCommand(rootCommand)

	return &CLI{
		rootCommand: rootCommand,
		command:     command,
	}
}

type CLI struct {
	rootCommand *cobra.Command
	command     *command.Command
}

func (cli *CLI) Register() {
	cli.command.Register()
}

func (cli *CLI) Start() {
	cli.rootCommand.Execute()
}
