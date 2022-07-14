package main

import (
	"fmt"
)

func closureExample() func() {
	x := 10
	funcao := func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimirX := closureExample()
	imprimirX()
}
