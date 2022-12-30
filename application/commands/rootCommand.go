package commands

import (
	"github.com/samber/do"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	CobraCommand *cobra.Command
}

func NewRootCommand(i *do.Injector) (*RootCommand, error) {
	cmd := &cobra.Command{
		Short: "RootCommand",
		Long:  `RootCommand`,
	}

	// Register sub commands
	cmd.AddCommand(do.MustInvoke[*Command1](i).CobraCommand)

	return &RootCommand{
		CobraCommand: cmd,
	}, nil
}
