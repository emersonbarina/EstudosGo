package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {

	}).Methods(http.MethodPost)

	fmt.Println("Escutando na porta 5001")
	log.Fatal(http.ListenAndServe(":5001", router))

}
