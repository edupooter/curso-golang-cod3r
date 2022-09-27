package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	fmt.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
