package server

import (
	"bd/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    int32  `json:"id"`
	Name  string `json:"nome"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte("Erro, falha ao ler o corpo"))
		return
	}

	var user user
	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("error"))
		return
	}

	db, err := banco.Connect()

	if err != nil {
		w.Write([]byte("erro no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into users (name, email) values (?, ?)")

	if err != nil {
		w.Write([]byte("erro no banco de dados"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)

	if erro != nil {
		w.Write([]byte("1"))
		return
	}
	id, err := insert.LastInsertId()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso id: %d", id)))
}
