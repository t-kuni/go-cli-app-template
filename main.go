package main

import (
	"fmt"
	"github.com/sarulabs/di"
)

func main() {
	generator := App().Get("quotation-generator").(QuotationGenerator)

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
				return RandomQuotationGenerator{
					clocker: App().Get("clocker").(Clocker),
				}, nil
			},
		},
		{
			Name: "clocker",
			Build: func(ctn di.Container) (interface{}, error) {
				return Clock{}, nil
			},
		},
	}...)

	postHook(builder)

	return builder.Build()
}
