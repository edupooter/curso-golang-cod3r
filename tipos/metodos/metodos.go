package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p Pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *Pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := Pessoa{
		nome:      "João",
		sobrenome: "Santos",
	}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Gabriela Anjo")
	fmt.Println(p1.getNomeCompleto())
}
