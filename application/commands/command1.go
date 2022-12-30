package commands

import (
	"github.com/samber/do"
	"github.com/spf13/cobra"
	"github.com/t-kuni/go-cli-app-template/domain/usecase"
)

type Command1 struct {
	CobraCommand *cobra.Command
}

func NewCommand1(i *do.Injector) (*Command1, error) {
	useCase := do.MustInvoke[*usecase.Command1UseCase](i)

	cmd := &cobra.Command{
		Use:   "command1",
		Short: "Command1",
		Long:  `Command1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return useCase.Execute()
		},
	}

	cmd.Flags().StringP("message", "m", "(default message)", "Message")

	return &Command1{
		CobraCommand: cmd,
	}, nil
}
