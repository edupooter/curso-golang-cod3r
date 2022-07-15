package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	// passagem por valor
	inc1(n)
	fmt.Println(n)

	// obtendo o endereço da variável e passagem por referência
	inc2(&n)
	fmt.Println(n)
}
