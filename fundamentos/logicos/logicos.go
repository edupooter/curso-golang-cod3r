package main

import "fmt"

func comprar(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTV50 := trabalho1 && trabalho2
	comprarTV32 := trabalho1 != trabalho2
	comprarSorvete := trabalho1 || trabalho2

	return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("tv50: %t\n", tv50)
	fmt.Printf("tv32: %t\n", tv32)
	fmt.Printf("sorvete: %t\n", sorvete)
	fmt.Printf("saudável: %t\n", !sorvete)
}
