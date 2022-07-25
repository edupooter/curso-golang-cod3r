package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// Dados enviados para o canal (escrita)
	ch <- 1

	// Recebendo dados do canal (leitura)
	<-ch

	ch <- 2

	fmt.Println(<-ch)
}
