package main

import "fmt"

type Curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type Dinamico interface{}
	var coisa2 Dinamico = "Olá"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = Curso{"Golang: Explorando a linguagem do Google"}
	fmt.Println(coisa2)
}
