package main

// インタフェース
type QuotationGenerator interface {
	GenerateQuotation() string
}

// 構造体
type SimpleQuotationGenerator struct {
}

// 構造体に関数を追加
func (generator SimpleQuotationGenerator) GenerateQuotation() string {
	return "貧すれば鈍する"
}

// 構造体
type EnglishQuotationGenerator struct {
}

// 構造体に関数を追加
func (generator EnglishQuotationGenerator) GenerateQuotation() string {
	return "poverty dulls the wit"
}
