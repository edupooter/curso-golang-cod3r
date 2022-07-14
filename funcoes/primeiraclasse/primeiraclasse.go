package main

import "fmt"

var somar = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(somar(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(4, 2))
}
