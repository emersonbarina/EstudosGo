package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("carregar página usuários!"))
}

func main() {
	// Conteitos básicos
	// HTTP é um protoco de comunicação - Base de comunicação web
	// Cliente - Servidor
	// Request - Response - Troca de mensagens
	// Rotas - identificar o tipo de mensagem
	// URI - identificador de recursos
	// Método - GET, POST, PUT, DELETE

	// Criando uma rota
	// testar em browse : http://localhost:5001/home
	http.HandleFunc("/home", home)

	http.HandleFunc("/user", user)

	// Criando um servidor para testes
	log.Fatal(http.ListenAndServe(":5001", nil))

}
