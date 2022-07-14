package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")

	for i, a := range aprovados {
		fmt.Printf("%d: %s\n", i+1, a)
	}
}

func main() {
	aprovados := []string{"Alberto", "Bernardo", "Carlos"}
	imprimirAprovados(aprovados...)
}
