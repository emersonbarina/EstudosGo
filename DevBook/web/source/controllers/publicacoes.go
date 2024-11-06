package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/source/config"
	"webapp/source/requests"
	"webapp/source/responses"

	"github.com/gorilla/mux"
)

// CriarPublicacao chama a API para cadastrar uma publicacao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requests.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
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

// CurtirPublicacao chapa API para adicionar curtidas de uma publicação
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	// fmt.Println(parametros)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoID"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.APIURL, publicacaoID)
	response, erro := requests.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
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
