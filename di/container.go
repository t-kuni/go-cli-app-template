package di

import (
	"github.com/samber/do"
	command2 "github.com/t-kuni/go-cli-app-template/application/commands"
	"github.com/t-kuni/go-cli-app-template/domain/services"
	"github.com/t-kuni/go-cli-app-template/domain/usecase"
	"github.com/t-kuni/go-cli-app-template/infrastructure/external"
	"github.com/t-kuni/go-cli-app-template/infrastructure/system"
)

func NewContainer() *do.Injector {
	injector := do.New()

	// Presentation
	do.Provide(injector, command2.NewRootCommand)
	do.Provide(injector, command2.NewCommand1)

	// Infrastructure
	do.Provide(injector, system.NewTimer)
	do.Provide(injector, system.NewStdio)
	do.Provide(injector, external.NewTodoClient)

	// UseCase
	do.Provide(injector, usecase.NewCommand1UseCase)

	// Domain
	do.Provide(injector, services.NewRandomQuotationGenerator)

	return injector
}
