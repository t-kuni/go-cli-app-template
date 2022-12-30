package di

import (
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/domain/services"
	"github.com/t-kuni/go-cli-app-template/domain/usecase"
	"github.com/t-kuni/go-cli-app-template/infrastructure/system"
	"github.com/t-kuni/go-cli-app-template/presentation/command"
)

func NewContainer() *do.Injector {
	injector := do.New()

	// Presentation
	do.Provide(injector, command.NewRootCommand)
	do.Provide(injector, command.NewCommand1)

	// Infrastructure
	do.Provide(injector, system.NewTimer)
	do.Provide(injector, system.NewStdio)

	// UseCase
	do.Provide(injector, usecase.NewCommand1UseCase)

	// Domain
	do.Provide(injector, services.NewRandomQuotationGenerator)

	return injector
}
