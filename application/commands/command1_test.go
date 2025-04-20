package commands_test

import (
	"context"
	"testing"
	"time"

	"github.com/rotisserie/eris"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/application/commands"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/external"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/system"
	"github.com/t-kuni/go-cli-app-template/domain/model"
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
			mock.EXPECT().Printf("Todo title： %s\n", "delectus aut autem").Times(1)
			do.OverrideValue[system.IStdio](cont.DI, mock)
		}
		{
			mock := external.NewMockITodoClient(cont.MockCtrl)
			todo := &model.Todo{
				UserID:    1,
				ID:        1,
				Title:     "delectus aut autem",
				Completed: false,
			}
			mock.
				EXPECT().
				FetchTodo(1).
				Return(todo, nil)
			do.OverrideValue[external.ITodoClient](cont.DI, mock)
		}

		//
		// Execute
		//
		testee := do.MustInvoke[*commands.RootCommand](cont.DI)
		cmd := testee.CobraCommand
		cmd.SetArgs([]string{"command1"})
		err := cmd.ExecuteContext(context.Background())

		// Assert
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Should handle error with stack trace", func(t *testing.T) {
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
		{
			// TodoClientのモックを作成し、エラーを返すように設定
			mock := external.NewMockITodoClient(cont.MockCtrl)
			expectedError := eris.New("テスト用のエラー")
			mock.
				EXPECT().
				FetchTodo(1).
				Return((*model.Todo)(nil), expectedError)
			do.OverrideValue[external.ITodoClient](cont.DI, mock)
		}
		// Loggerのモックは必要ありません。main.goでのみ使用されるため、テストでは呼ばれません。

		//
		// Execute
		//
		testee := do.MustInvoke[*commands.RootCommand](cont.DI)
		cmd := testee.CobraCommand
		cmd.SetArgs([]string{"command1"})
		err := cmd.ExecuteContext(context.Background())

		// Assert
		if err == nil {
			t.Fatal("エラーが発生するはずですが、nilが返されました")
		}
	})
}
