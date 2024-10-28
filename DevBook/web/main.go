package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/source/router"
)

func main() {
	fmt.Println("Rodando WebApp!")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3001", r))
}
