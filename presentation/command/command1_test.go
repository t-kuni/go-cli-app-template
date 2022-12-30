package command_test

import (
	"context"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/system"
	"github.com/t-kuni/go-cli-app-template/presentation/command"
	"testing"
	"time"
)

func TestCommand1(t *testing.T) {
	t.Run("Should execute normally", func(t *testing.T) {
		//
		// Prepare
		//
		cont := beforeEach(t)
		defer afterEach(cont)

		{
			mock := system.NewMockITimer(cont.MockCtrl)
			mock.
				EXPECT().
				Now().
				Return(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
			do.OverrideValue[system.ITimer](cont.DI, mock)
		}
		{
			mock := system.NewMockIStdio(cont.MockCtrl)
			mock.EXPECT().Printf("App name： %s\n", "app-name").Times(1)
			mock.EXPECT().Printf("Quotation： %s\n", "money comes and goes; money goes around and around").Times(1)
			do.OverrideValue[system.IStdio](cont.DI, mock)
		}

		//
		// Execute
		//
		testee := do.MustInvoke[*command.RootCommand](cont.DI)
		cmd := testee.CobraCommand
		cmd.SetArgs([]string{"command1"})
		err := cmd.ExecuteContext(context.Background())

		// Assert
		if err != nil {
			t.Fatal(err)
		}
	})
}
