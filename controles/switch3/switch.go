package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float64, float32:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	case bool:
		return "bool"
	case complex128, complex64:
		return "complexo"
	default:
		return "tipo inválido"
	}
}

func main() {
	fmt.Println(tipo(54))
	fmt.Println(tipo(2.3))
	fmt.Println(tipo("Hello"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
	fmt.Println(tipo(1 + 2i))
	// This function doesn't recognize 'time' types
	fmt.Println(tipo(time.Now()))
}
