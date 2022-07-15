package main

import "fmt"

type Imprimivel interface {
	toString() string
}

type Pessoa struct {
	nome      string
	sobrenome string
}

type Produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (pessoa Pessoa) toString() string {
	return pessoa.nome + " " + pessoa.sobrenome
}

func (produto Produto) toString() string {
	return fmt.Sprintf("%s - R$%.2f", produto.nome, produto.preco)
}

func imprimir(x Imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var objeto1 Imprimivel = Pessoa{"Eduardo", "Reis"}
	fmt.Println(objeto1.toString())

	objeto1 = Produto{"Notebook", 3999.80}
	imprimir(objeto1)

	imprimir(Produto{"Cadeira", 299.11})
}
