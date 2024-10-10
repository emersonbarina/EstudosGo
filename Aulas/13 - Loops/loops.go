package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0
	// fmt.Println("incrementando i")
	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println("Concluído!")

	// fmt.Println("incrementando J")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("...", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println("Concluído!")

	nomes := [3]string{"João", "José", "Maria"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}
	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Emerson",
		"sobrenome": "Barina",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		// loop infinito
		fmt.Println("Executando...")
		time.Sleep(time.Second)
	}

}
