package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/application/commands"
	"github.com/t-kuni/go-cli-app-template/di"
	"os"
)

func main() {
	ctx := context.Background()

	godotenv.Load()

	container := di.NewContainer()
	defer container.Shutdown()

	cmd := do.MustInvoke[*commands.RootCommand](container)
	err := cmd.CobraCommand.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}
