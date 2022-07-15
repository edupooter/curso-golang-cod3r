package main

import "fmt"

type Esportivo interface {
	ligarTurbo()
}

type Luxuoso interface {
	fazerBaliza()
}

type EsportivoLuxuoso interface {
	Esportivo
	Luxuoso
}

type BMW7 struct{}

func (b BMW7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b BMW7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b EsportivoLuxuoso = BMW7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
