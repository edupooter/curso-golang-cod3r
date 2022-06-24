package main

import (
	"fmt"
)

func main() {
	x := 3.141592653589793

	stringx := fmt.Sprint(x)

	fmt.Println("x value is:", stringx)

	fmt.Printf("x value is: %.2f", x)

	a, b, c, d := 1, 1.999, true, "false"

	// Exemplo com "%v", que aceita diferentes tipos de valores
	fmt.Printf("\nvalores: %v, %.1f, %v, %v", a, b, c, d)
}
