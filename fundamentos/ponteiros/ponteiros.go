package main

import "fmt"

func main() {
	i := 1

	// Golang não permite aritmética de ponteiros
	var p *int = nil

	// atribui o endereço da variável i para a variável p (do tipo ponteiro)
	p = &i

	// desreferenciação (recupera o valor)
	*p++

	i++

	fmt.Println(p, *p, i, &i)
}
