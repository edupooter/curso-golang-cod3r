package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"João Silva":      1234.56,
		"Maria Santos":    3521.49,
		"Gabriela Vieira": 2004.19,
	}

	funcsESalarios["Pedro Junior"] = 1088.17

	delete(funcsESalarios, "Zé Ninguém")

	for funcionario, salario := range funcsESalarios {
		fmt.Printf("%v: %v\n", funcionario, salario)
	}
}
