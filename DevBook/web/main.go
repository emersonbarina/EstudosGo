package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/source/config"
	"webapp/source/cookies"
	"webapp/source/router"
	"webapp/source/utils"
)

// func init() {
// 	hasKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hasKey)

// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blockKey)
// }

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
