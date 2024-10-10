package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle - Case")
	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("é menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
		fmt.Println(outroNumero)
	} else if outroNumero == 0 {
		fmt.Println("Número é igual a zero")
		fmt.Println(outroNumero)
	} else {
		fmt.Println("Número é menor que zero")
		fmt.Println(outroNumero)
	}
}
