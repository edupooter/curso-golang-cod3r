package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.147)/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES (?)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	idLast, _ := res.LastInsertId()

	fmt.Println("Último id inserido:", idLast)

	linhas, _ := res.RowsAffected()

	fmt.Println("Linhas afetadas no último comando:", linhas)
}
