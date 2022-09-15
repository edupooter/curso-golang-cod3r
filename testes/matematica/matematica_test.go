package matematica

import "testing"

const erroPadrao = "Valor esperado era %v e valor resultado foi %v"

func TestObterMedia(t *testing.T) {
	valorEsperado := 7.13

	valorResultado := Media(7.2, 9.9, 6.1, 5.3)

	if valorEsperado != valorResultado {
		t.Errorf(erroPadrao, valorEsperado, valorResultado)
	}
}
