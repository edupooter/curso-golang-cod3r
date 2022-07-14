package main

import "fmt"

func calcularMedia(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f\n", calcularMedia(5, 9.234656, 123))
}
