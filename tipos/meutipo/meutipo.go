package main

import "fmt"

type Nota float64

func (n Nota) verificarNotaEntre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func converterNotaParaConceito(n Nota) string {
	switch {
	case n.verificarNotaEntre(9, 10):
		return "A"
	case n.verificarNotaEntre(7, 8.99):
		return "B"
	case n.verificarNotaEntre(5, 7.99):
		return "C"
	case n.verificarNotaEntre(3, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	conceito := converterNotaParaConceito(7.5)
	fmt.Println(conceito)
}
