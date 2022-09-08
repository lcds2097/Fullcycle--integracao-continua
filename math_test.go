package main 

import "testing"

func TestSoma(t *testing.T) {
	total := Somar(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}


func TestSubtracao(t *testing.T) {
	total := Subtrair(15, 15)

	if total != 0 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 0)
	}
}


func TestMultiplicacao(t *testing.T) {
	total := Multiplicar(15, 2)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}


func TestDivisao(t *testing.T) {
	total := Dividir(10, 2)

	if total != 5 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 5)
	}
}