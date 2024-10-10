package main

import "fmt"

func main() {
	canal := make(chan string, 2) // buffer = 2
	canal <- "Olá mundo!"
	canal <- "Programando em Go!"

	mensagem1 := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem1, mensagem2)

}
