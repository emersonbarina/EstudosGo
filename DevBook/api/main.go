package main

import (
	"api/source/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API...")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5001", r))

}
