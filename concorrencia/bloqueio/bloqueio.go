package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	// Operação bloqueante, pois o canal não possui buffer
	c <- 1

	fmt.Println("Impresso após lido")
}

func main() {
	c := make(chan int)
	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
}
