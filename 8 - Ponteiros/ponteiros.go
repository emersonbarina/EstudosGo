package main

import "fmt"

func main() {
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var2++
	fmt.Println(var1, var2)

	// Ponteiro é uma referência de memória
	var var3 int
	var ponteiro *int

	var3 = 1
	ponteiro = &var3
	fmt.Println(var3, ponteiro)

	fmt.Println(var3, *ponteiro) // desreferenciação

	var3 = 100
	fmt.Println(var3, ponteiro)

	fmt.Println(var3, *ponteiro) // desreferenciação
}
