package main

import "fmt"

type Item struct {
	produtoId int
	qtde      float64
	preco     float64
}

type Pedido struct {
	userId int
	itens  []Item
}

func (p Pedido) calcularPedido() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * item.qtde
	}

	return total
}

func main() {
	pedido := Pedido{
		userId: 123,
		itens: []Item{
			{produtoId: 9, qtde: 2.3, preco: 3.50},
			{produtoId: 7, qtde: 1, preco: 8.10},
			{produtoId: 5, qtde: 30.55, preco: 1.99},
		},
	}

	fmt.Printf("Valor total: R$%.2f\n", pedido.calcularPedido())
}
