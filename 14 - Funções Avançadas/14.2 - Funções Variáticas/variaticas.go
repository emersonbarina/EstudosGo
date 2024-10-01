package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	soma := 0
	for _, valor := range numeros {
		fmt.Println(soma, valor)
		soma += valor
	}
	return soma
}

func escrever(texto string, numeros ...int) {
	for _, valor := range numeros {
		fmt.Println(texto, valor)
	}
}
func main() {
	total := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 200, 11, 15, 16)
	fmt.Println(total)
	escrever("Teste", 10, 10, 23, 2, 3, 4, 5, 6)
}
