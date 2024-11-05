package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/source/config"
	"webapp/source/models"
	"webapp/source/requests"
	"webapp/source/responses"
	"webapp/source/utils"
)

// CarregarTelaDeLogin vai reenderizar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario vai carregar a página de cadastro de usuários
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal vai carregar a página home com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)

	// response, erro := http.Get(url)
	response, erro := requests.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	// fmt.Println(response.StatusCode, erro)
	var publicacoes []models.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	// fmt.Println(publicacoes)

	utils.ExecutarTemplate(w, "home.html", publicacoes)
	// utils.ExecutarTemplate(w, "home.html", struct {
	// 	Publicacoes []models.Publicacao
	// 	OutroCampo  string
	// }{
	// 	Publicacoes: publicacoes,
	// 	OutroCampo:  "Valor qualquer",
	// })
}
