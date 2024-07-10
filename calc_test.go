package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	//arrange
	teste := somar(3, 2, 1)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}

func TestSoma2(t *testing.T) {
	//arrange
	teste := somar(3, 2, 1)

	//act
	resultado := 7

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}

func TestSubtrair(t *testing.T) {
	//arrange
	teste := subtrair(2, 2)

	//act
	resultado := 0

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}
func TestSubtrair2(t *testing.T) {
	//arrange
	teste := subtrair(2, 2)

	//act
	resultado := -1

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}

func TestMultiplicar(t *testing.T) {
	//arrange
	teste := multiplicar(3, 2)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}

func TestMultiplicar2(t *testing.T) {
	//arrange
	teste := somar(3, 2)

	//act
	resultado := 10

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}

func TestDividir(t *testing.T) {
	//arrange
	teste := dividir(2, 6)

	//act
	resultado := 3

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}

func TestDividir2(t *testing.T) {
	//arrange
	teste := dividir(2, 6)

	//act
	resultado := 4

	//assert
	if teste != resultado {
		t.Error("Valor Esperado:", resultado, "Valor Retornado:", teste)
	}
}
