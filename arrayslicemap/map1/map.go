package main

import "fmt"

func main() {
	// mapas devem ser inicializados antes do uso
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[1111111111] = "Maria"
	aprovados[2222233333] = "Pedro"
	aprovados[9985666666] = "Ana"

	for id, name := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", name, id)
	}

	delete(aprovados, 1111111111)

	for id, name := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", name, id)
	}
}
