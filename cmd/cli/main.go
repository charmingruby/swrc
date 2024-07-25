package main

import (
	"github.com/charmingruby/swrc/internal/infra/cli"
	"github.com/spf13/cobra"
)

func main() {
	var rootCommand = &cobra.Command{}
	cli := cli.NewCLI(rootCommand)
	cli.Register()
	cli.Start()
}
