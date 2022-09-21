package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "wsl_root:123456@/@192.168.0.147:3306")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists cursogo")
}
