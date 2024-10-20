package main

import (
	"api/source/config"
	"api/source/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	// fmt.Println(config.StringConexaoBanco)
	fmt.Printf("Rodando API na porta %d...", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
