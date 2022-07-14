package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  1265.55,
			"Gustavo Pereira": 6246.56,
		},
		"J": {
			"Jenna Silva": 5326.00,
			"João Maria":  1026.5,
		},
		"P": {
			"Pedro Junior": 1209.99,
		},
	}

	delete(funcsPorLetra["J"], "João Maria")

	for letra, funcionario := range funcsPorLetra {
		fmt.Println("Letra:", letra)
		for i, salario := range funcionario {
			fmt.Println(i, salario)
		}
	}
}
