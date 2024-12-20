package main

import (
	"api/source/config"
	"api/source/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	// Gerando um chave randomica para o Secret_key
// 	chave := make([]byte, 64)

// 	if _, erro := rand.Read((chave)); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Carregar()
	// fmt.Println(config.StringConexaoBanco)
	fmt.Printf("Rodando API na porta %d...", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
