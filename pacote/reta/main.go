package main

import "fmt"

func main() {
	p1 := Ponto{2, 2}
	p2 := Ponto{4, 9}

	fmt.Println(calcularCatetos(p1, p2))
	fmt.Println(CalcularDistancia(p1, p2))
}
