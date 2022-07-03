package usecase

import (
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/domain/infrastructure/system"
	"github.com/t-kuni/go-cli-app-skeleton/domain/services"
	"os"
)

type Command1UseCase struct {
	quotationGenerator services.IQuotationGenerator
	stdio              system.IStdio
}

func NewCommand1UseCase(i *do.Injector) (*Command1UseCase, error) {
	return &Command1UseCase{
		quotationGenerator: do.MustInvoke[services.IQuotationGenerator](i),
		stdio:              do.MustInvoke[system.IStdio](i),
	}, nil
}

func (u Command1UseCase) Execute() error {
	appName := os.Getenv("APP_NAME")

	u.stdio.Printf("App name： %s\n", appName)

	u.stdio.Printf("Quotation： %s\n", u.quotationGenerator.GenerateQuotation())
	return nil
}
