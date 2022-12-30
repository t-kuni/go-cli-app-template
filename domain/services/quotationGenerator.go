package services

import (
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/system"
)

type IQuotationGenerator interface {
	GenerateQuotation() string
}

type RandomQuotationGenerator struct {
	Timer system.ITimer
}

func NewRandomQuotationGenerator(i *do.Injector) (IQuotationGenerator, error) {
	return RandomQuotationGenerator{
		Timer: do.MustInvoke[system.ITimer](i),
	}, nil
}

func (s RandomQuotationGenerator) GenerateQuotation() string {
	t := s.Timer.Now()
	if t.Second()%2 == 1 {
		return "Time is money"
	} else {
		return "money comes and goes; money goes around and around"
	}
}
