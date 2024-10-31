package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/responses"
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
		// log.Fatal(erro)
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	// fmt.Println(usuario)
	// fmt.Println(bytes.NewBuffer(usuario))

	response, erro := http.Post("http://localhost:5001/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		// log.Fatal(erro)
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
