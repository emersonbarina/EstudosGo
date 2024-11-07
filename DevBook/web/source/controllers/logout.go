package controllers

import (
	"net/http"
	"webapp/source/cookies"
)

// FazerLogout remove os dados de autenticação do browse do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
