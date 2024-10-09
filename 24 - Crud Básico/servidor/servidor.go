package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario inser um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Falha ao converter o usuário para struct"))
		return
	}

	// fmt.Println(usuario)
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Falha ao criar statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Falha ao inserir dados"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Falha ao obter o id inserido!"))
		return
	}

	// STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso. Id: %d", idInserido)))
}
