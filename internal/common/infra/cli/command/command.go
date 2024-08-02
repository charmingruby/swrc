package command

import (
	"github.com/charmingruby/swrc/internal/common/infra/transport/grpc/client"
	"github.com/spf13/cobra"
)

func NewCommand(rootCommand *cobra.Command, client *client.CommonClientHandler) *Command {
	return &Command{
		rootCommand: rootCommand,
		client:      client,
	}
}

type Command struct {
	rootCommand *cobra.Command
	client      *client.CommonClientHandler
}

func (cmd *Command) Register() {
	cmd.rootCommand.AddCommand(cmd.newHealthCheckCommand())
}
