package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	a := example(1, 2)
	if a != 3 {
		t.Errorf("失敗")
	}
}
