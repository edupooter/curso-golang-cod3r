package main

import (
	"fmt"
	"math"
	"reflect"
)

func maxInts(nums ...int) int {
	maxNum := -int(^uint(0)>>1) - 1
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func main() {
	// inteiros
	fmt.Println(1, -2, 1000000)
	fmt.Println("Literal inteiro:", reflect.TypeOf(1), reflect.TypeOf(-2), reflect.TypeOf(1000000))

	var b byte = 255

	fmt.Println(reflect.TypeOf(b))

	// valor inteiro máximo
	inteiro_maximo := math.MaxInt64

	fmt.Println(inteiro_maximo)

	maior := maxInts(1, 5659, 262, -4654, 99999999999)

	fmt.Println(maior)

	// Concatenação de string
	primeira := "hello"
	segunda := "world"

	fmt.Println(primeira + segunda + "!")

	// String literal
	multiplas_linhas := `<html>
	<head>
	</head>
	<body>
		Hello
	</body>
	</html>`

	fmt.Println(len(multiplas_linhas))
}
