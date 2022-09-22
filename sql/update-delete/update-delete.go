package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.147)/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// UPDATE
	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Maria Eduarda", 1)
	stmt.Exec("Sharon Valeska", 2)

	// DELETE
	stmtDel, _ := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	stmtDel.Exec(3)
}
