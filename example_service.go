package main

type ExampleService interface {
	Exec() string
}

func Exec() string {
	return "called example service"
}
