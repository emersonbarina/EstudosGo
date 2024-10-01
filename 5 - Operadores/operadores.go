package main

import (
	"fmt"
)

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 20
	var numero2 int16 = 35
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	// Atribuição
	var variavel1 string = "String"
	var variavel2 string = "String 2"
	fmt.Println(variavel1, variavel2)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores unários
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 5
	fmt.Println(numero)
	numero *= 2
	fmt.Println(numero)
	numero /= 3
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)

	// Operador ternário (não existo no GO)
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor ou igual a 5"
	}
	fmt.Print(texto)

}
