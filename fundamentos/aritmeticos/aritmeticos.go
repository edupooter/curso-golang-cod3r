package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Adição => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Multiplicação => ", a*b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Módulo => ", a%b)

	// bitwise
	fmt.Println("Bitwise AND => ", a&b) // 11 & 10 = 10
	fmt.Println("Bitwise OR => ", a|b)  // 11 | 10 = 11
	fmt.Println("Bitwise XOR => ", a^b) // 11 ^ 10 = 01
	fmt.Println("Bitwise << => ", a<<b) // 11 << 10 = 12
	fmt.Println("Bitwise >> => ", a>>b) // 11 >> 10 = 0

	c := 3.0
	d := 2.0

	fmt.Println("Maior => ", math.Max(c, d))
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Exponenciação => ", math.Pow(c, d))
}
