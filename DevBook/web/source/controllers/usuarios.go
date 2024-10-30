package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CriarUsuario chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	// fmt.Println(usuario)
	// fmt.Println(bytes.NewBuffer(usuario))

	response, erro := http.Post("http://localhost:5001/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		log.Fatal(erro)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)

}
