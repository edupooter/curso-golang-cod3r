package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo, primeiro = p2, p1
	// retorno limpo
	return
}

func main() {
	x, y := trocar(3, 5)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 9)
	fmt.Println(segundo, primeiro)
}
