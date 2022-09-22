package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.147)/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios WHERE id > ?", 0)

	defer rows.Close()

	for rows.Next() {
		var u usuario
		id := &u.id
		nome := &u.nome
		rows.Scan(id, nome)
		fmt.Println(u)
	}
}
