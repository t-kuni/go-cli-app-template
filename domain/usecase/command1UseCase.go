package usecase

import (
	"fmt"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/domain/services"
	"os"
)

type Command1UseCase struct {
	quotationGenerator services.IQuotationGenerator
}

func NewCommand1UseCase(i *do.Injector) (*Command1UseCase, error) {
	return &Command1UseCase{
		quotationGenerator: do.MustInvoke[services.IQuotationGenerator](i),
	}, nil
}

func (u Command1UseCase) Execute() error {
	appName := os.Getenv("APP_NAME")
	fmt.Printf("App name： %s\n", appName)

	fmt.Printf("Quotation： %s\n", u.quotationGenerator.GenerateQuotation())
	return nil
}
