package area

import "math"

// Pi é a proporção númerica entre o perímetro e o diâmetro de um círculo
const Pi = 3.141592653589793

// CalcularAreaCirculo calcula a área de um círculo
func CalcularAreaCirculo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// CalcularAreaRetangulo calcula a área de um retângulo
func CalcularAreaRetangulo(base, altura float64) float64 {
	return base * altura
}

// Não é visível fora do pacote
func _CalcularAreaTrianguloEquilatero(base, altura float64) float64 {
	return base * altura / 2
}
