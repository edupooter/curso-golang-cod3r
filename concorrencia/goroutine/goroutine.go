package main

import (
	"fmt"
	"time"
)

func falar(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração: %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// falar("Denise", "Quer falar comigo?", 3)
	// falar("Jorge", "Somente depois de você.", 1)

	// go falar("Maria", "Eita!", 500)
	// go falar("João", "Opa...", 500)

	go falar("Marta", "Entendi.", 10)
	falar("José", "Parabéns!", 5)
}
