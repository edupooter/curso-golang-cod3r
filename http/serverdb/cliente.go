package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuário
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// Handler de Usuário serve para analisar request e delegar para a função correta
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorId(w, r, id)
	case r.Method == "GET":
		usuarioList(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Não encontrado")
	}
}

func usuarioPorId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.147)/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var u Usuario

	db.QueryRow("SELECT id, nome FROM usuarios WHERE id = ? LIMIT 1", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(json))
}

func usuarioList(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.0.147)/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios")

	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
