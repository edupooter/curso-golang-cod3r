package main

import "fmt"

type Carro struct {
	nome            string
	velocidadeAtual int
}

type Ferrari struct {
	Carro       // campos anônimos via composição
	turboLigado bool
}

func main() {
	f := Ferrari{}
	f.nome = "F-50"
	f.velocidadeAtual = 0
	f.turboLigado = false

	fmt.Printf("A Ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
}
