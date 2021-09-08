package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
	"github.com/t-kuni/go-cli-app-skeleton/internal"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appName := os.Getenv("APP_NAME")
	fmt.Printf("アプリ名： %s\n", appName)

	generator := App().Get("quotation-generator").(internal.QuotationGenerator)
	fmt.Printf("名言： %s\n", generator.GenerateQuotation())
}

var app di.Container

func App() di.Container {
	if app != nil {
		return app
	}

	app := buildContainer(func(builder *di.Builder) {})

	return app
}

func buildContainer(postHook func(builder *di.Builder)) di.Container {
	builder, _ := di.NewBuilder()

	builder.Add([]di.Def{
		{
			Name: "quotation-generator",
			Build: func(ctn di.Container) (interface{}, error) {
				return internal.RandomQuotationGenerator{
					Clocker: App().Get("clocker").(internal.Clocker),
				}, nil
			},
		},
		{
			Name: "clocker",
			Build: func(ctn di.Container) (interface{}, error) {
				return internal.Clock{}, nil
			},
		},
	}...)

	postHook(builder)

	return builder.Build()
}
