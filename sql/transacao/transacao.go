package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.147)/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) values(?, ?)")

	stmt.Exec(2000, "Carlos")
	stmt.Exec(2001, "Beatriz")

	query, err := stmt.Exec(1, "Tiago") // teste com chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	fmt.Println(query)

	tx.Commit()
}
