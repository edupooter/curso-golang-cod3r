package main

import (
	"fmt"

	"github.com/edupooter/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)

	return c
}

func main() {
	c := multiplexar(
		html.Titulo("https://github.com/", "https://www.cod3r.com.br"),
		html.Titulo("https://www.google.com/", "https://www.terra.com.br"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
