package command

import (
	"github.com/charmingruby/swrc/pkg/cli/statement"
	"github.com/spf13/cobra"
)

func (cmd *Command) newHealthCheckCommand() *cobra.Command {
	var testInput string

	newCmd := &cobra.Command{
		Use:   "check",
		Short: "CLI health check",
		Run: func(cmd *cobra.Command, args []string) {
			if testInput == "" {
				statement.BreakLineStatement("You must supply a test input.")
				return
			}

			statement.BreakLineStatement("CLI is OK.")
		},
	}

	newCmd.Flags().StringVarP(&testInput, "input", "i", "", "Test input")

	return newCmd
}
