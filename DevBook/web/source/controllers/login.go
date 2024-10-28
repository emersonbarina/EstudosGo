package controllers

import (
	"net/http"
	"webapp/source/utils"
)

// CarregarTelaDeLogin vai reenderizar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
