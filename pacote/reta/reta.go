package main

import (
	"math"
)

// Letra inicial maiúscula significa visibilidade pública (independe do arquivo)
// Letra inicial minúscula ou prefixo "_", visibilidade privada no pacote

// Ponto representa uma coordenada cartesiana
type Ponto struct {
	x float64
	y float64
}

func calcularCatetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// CalcularDistancia é responsável por calcular a distância linear entre dois pontos
func CalcularDistancia(a, b Ponto) float64 {
	cx, cy := calcularCatetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
