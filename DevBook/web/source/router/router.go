package router

import (
	"webapp/source/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar retorna um router com todas rotas geradas

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
