package middlewares

import (
	"api/source/auth"
	"api/source/responses"
	"log"
	"net/http"
)

// Logger escreve informacões da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica se  o usuário fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidarToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}
