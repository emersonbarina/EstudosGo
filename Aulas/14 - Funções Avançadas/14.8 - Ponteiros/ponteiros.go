package main

import "fmt"

func inverterSinal(numero int) int {
	// passando cópia
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	// passando a referência
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println("Numero", numero, "Invertido", numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("Invertido", novoNumero)

}
