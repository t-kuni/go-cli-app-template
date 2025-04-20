package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/application/commands"
	"github.com/t-kuni/go-cli-app-template/di"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/system"
)

func main() {
	ctx := context.Background()

	godotenv.Load()

	container := di.NewContainer()
	defer container.Shutdown()

	cmd := do.MustInvoke[*commands.RootCommand](container)
	err := cmd.CobraCommand.ExecuteContext(ctx)
	if err != nil {
		logger := do.MustInvoke[system.ILogger](container)
		logger.WithError(err).Error("エラーが発生しました")
		os.Exit(1)
	}
}
