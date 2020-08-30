package internal

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

// 構造体
type RandomQuotationGenerator struct {
	Clocker Clocker
}

// 構造体に関数を追加
func (s RandomQuotationGenerator) GenerateQuotation() string {
	t := s.Clocker.Now()
	if t.Second()%2 == 1 {
		return "時は金なり"
	} else {
		return "金は天下の回り物"
	}
}
