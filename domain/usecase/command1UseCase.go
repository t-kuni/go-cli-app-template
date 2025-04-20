package usecase

import (
	"os"

	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/api"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/system"
	"github.com/t-kuni/go-cli-app-template/domain/services"
)

type Command1UseCase struct {
	quotationGenerator services.IQuotationGenerator
	stdio              system.IStdio
	todoClient         api.ITodoClient
}

func NewCommand1UseCase(i *do.Injector) (*Command1UseCase, error) {
	return &Command1UseCase{
		quotationGenerator: do.MustInvoke[services.IQuotationGenerator](i),
		stdio:              do.MustInvoke[system.IStdio](i),
		todoClient:         do.MustInvoke[api.ITodoClient](i),
	}, nil
}

func (u Command1UseCase) Execute() error {
	appName := os.Getenv("APP_NAME")

	u.stdio.Printf("App name： %s\n", appName)

	u.stdio.Printf("Quotation： %s\n", u.quotationGenerator.GenerateQuotation())

	// JSONPlaceholderからTodoを取得
	todo, err := u.todoClient.FetchTodo(1)
	if err != nil {
		return err
	}

	// Todoのタイトルを出力
	u.stdio.Printf("Todo title： %s\n", todo.Title)

	return nil
}
