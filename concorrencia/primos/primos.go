package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primes(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrime(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 1000)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 15)

	go primes(cap(c), c)

	for prime := range c {
		fmt.Printf("%d ", prime)
	}

	fmt.Println("Fim")
}
