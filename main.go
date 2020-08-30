package main

import (
	"fmt"
	"github.com/sarulabs/di"
)

func main() {
	app := createApp()

	generator := app.Get("quotation-generator").(QuotationGenerator)

	fmt.Printf("名言： %s\n", generator.GenerateQuotation())
	fmt.Printf("%d", example(1, 2))
}

func example(a int, b int) int {
	return a + b
}

func createApp() di.Container {
	builder, _ := di.NewBuilder()

	builder.Add([]di.Def{
		{
			Name: "quotation-generator",
			Build: func(ctn di.Container) (interface{}, error) {
				return SimpleQuotationGenerator{}, nil
			},
		},
	}...)

	return builder.Build()
}
