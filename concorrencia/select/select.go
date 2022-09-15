package main

import (
	"fmt"
	"time"

	"github.com/edupooter/html"
)

func faster(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// Select: estrutura de controle espefícia para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderem"
		// default:
		// 	return "Sem resposta"
	}
}

func main() {
	campeao := faster(
		"https://github.com/",
		"https://www.terra.com.br/",
		"https://www.google.com.br/",
	)

	fmt.Println(campeao)
}
