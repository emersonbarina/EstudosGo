package router

import "github.com/gorilla/mux"

// Gerar retorna um router com todas rotas geradas

func Gerar() *mux.Router {
	return mux.NewRouter()
}
