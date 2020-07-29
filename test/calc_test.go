package test

import (
	"../calc"
	"testing"
)

func TestSum(t *testing.T) {
	total := calc.Sum{X: 5, Y: 5}
	if total.Calc() != 10 {
		t.Errorf("Valor da Soma esta incorreto onde %d + %d = %d, valor esperado: %d.", total.X, total.Y, total.Calc(), 10)
	}
}