package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/source/router"
	"webapp/source/utils"
)

func main() {
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Escutando na porta 3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}
