package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	// convertendo para json
	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	// retornou um slice de unint8 = []unit8
	fmt.Println(cachorroJson)

	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"nome": "Toby",
		"reca": "Poodle",
	}
	cachorro2JSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	// retornou um slice de unint8 = []unit8
	fmt.Println(cachorro2JSON)

	fmt.Println(bytes.NewBuffer(cachorro2JSON))

}
