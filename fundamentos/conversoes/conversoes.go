package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado ao converter de string
	fmt.Println("Exemplo " + string(97))

	// int para string
	fmt.Println("Para string " + strconv.Itoa(123))
	fmt.Println("Para string " + strconv.FormatInt(123, 10))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 120)

	// conversão de boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
