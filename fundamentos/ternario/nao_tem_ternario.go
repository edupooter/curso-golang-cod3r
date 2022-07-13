package main

import "fmt"

// Não tem ternário
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	resultado := obterResultado(5.9)
	fmt.Printf("Resultado: %v\n", resultado)
}
