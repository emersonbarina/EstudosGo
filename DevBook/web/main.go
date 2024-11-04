package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/source/config"
	"webapp/source/router"
	"webapp/source/utils"
)

func main() {
	config.Carregar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
