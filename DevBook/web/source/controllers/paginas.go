package controllers

import (
	"fmt"
	"net/http"
	"webapp/source/config"
	"webapp/source/requests"
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

	fmt.Println(response.StatusCode, erro)

	utils.ExecutarTemplate(w, "home.html", nil)
}
