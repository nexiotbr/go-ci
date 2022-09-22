package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma inválido: Resultado: %d, valor esperado %d", total, 30)
	}
}

func TestSubstracao(t *testing.T) {
	total := Substracao(15, 15)

	if total != 0 {
		t.Errorf("Resultado da substação inválido: Resultado: %d, valor esperado %d", total, 0)
	}
}

func TestMultiplicacao(t *testing.T) {
	total := Multiplicacao(8, 5)

	if total != 40 {
		t.Errorf("Resultado da multiplicação inválido: Resultado: %d, valor esperado %d", total, 40)
	}
}
