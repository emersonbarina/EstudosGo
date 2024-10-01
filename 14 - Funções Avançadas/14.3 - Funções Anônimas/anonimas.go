package main

import "fmt"

func main() {

	func() {
		fmt.Println("Funções Anônimas...")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Funções Anônimas...")

	retorno := func(texto string) string {
		return fmt.Sprintf("Teste -> %s %d", texto, 10)
	}("Funções Anônimas...")
	fmt.Println(retorno)
}
