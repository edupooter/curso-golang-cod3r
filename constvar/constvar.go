package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.141592653589793
	var raio = 3.2

	area := PI * m.Pow(raio, 2)

	fmt.Println("área igual a:", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Printf("constantes a = %d, b = %d, c = %d, d = %d\n", a, b, c, d)

	e, f := true, "string"

	fmt.Println(e, f)

	g, h := 0.1, 0.2

	fmt.Println("g + h = ", g+h)
}
