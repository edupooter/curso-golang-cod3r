package main

import (
	"encoding/json"
	"fmt"
)

type Produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := Produto{1, "Notebook", 1899.80, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 Produto
	jsonString := `{"id":2,"nome":"Cadeira","preco":690.23,"tags":["Importado","Móvel"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[0])
}
