package internal

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateQuotation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	loc, _ := time.LoadLocation("Asia/Tokyo")
	clocker := NewMockClocker(ctrl)
	clocker.EXPECT().Now().Return(time.Date(2020, 1, 1, 0, 0, 0, 0, loc))

	actual := RandomQuotationGenerator{Clocker: clocker}

	assert.Equal(t, "金は天下の回り物", actual.GenerateQuotation())
}
