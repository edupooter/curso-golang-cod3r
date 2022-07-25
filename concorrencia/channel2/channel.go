package main

import (
	"fmt"
	"time"
)

// Canal é o meio de comunicação entre goroutines
// Channel é um tipo, first class citizen

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- base * 2

	time.Sleep(time.Second)
	c <- base * 3

	time.Sleep(time.Second * 3)
	c <- base * 4
}

func main() {
	c := make(chan int)

	go doisTresQuatroVezes(2, c)

	// Recebendo dados do canal
	a, b := <-c, <-c

	fmt.Println(a, b)

	fmt.Println(<-c)
}
