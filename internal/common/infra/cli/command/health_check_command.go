package command

import (
	"github.com/charmingruby/swrc/pkg/cli/statement"
	"github.com/charmingruby/swrc/proto/pb"
	"github.com/spf13/cobra"
)

func (cmd *Command) newHealthCheckCommand() *cobra.Command {
	var greeting string

	newCmd := &cobra.Command{
		Use:   "hc",
		Short: "Health check",
		Run: func(c *cobra.Command, args []string) {
			if greeting == "" {
				statement.BreakLineStatement("You must supply a greeting.")
				return
			}

			_, err := cmd.client.HealthCheck(c.Context(), &pb.PingMessage{Greeting: greeting})
			if err != nil {
				statement.BreakLineStatement("Unhealthy server.")
				return
			}

			statement.BreakLineStatement("Ready to work.")
		},
	}

	newCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Greeting")

	return newCmd
}
