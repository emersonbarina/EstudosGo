package main

import (
	"errors"
	"fmt"
)

func main() {
	// 4 tipo de inteiros int8, int16, int32, int64 e o int que usa a arquitetura do computador, aceita negativo
	var numero int = 10000000000000
	fmt.Println(numero)

	// uint - int sem sinal, não aceita negativo
	var numero2 uint = 100000
	fmt.Println(numero2)

	// Alias
	// INT32 = RUNE
	// UINT8 = BYTE
	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	//float Números Reais
	var numeroReal1 float32 = 123.
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.
	fmt.Println(numeroReal2)

	numeroReal3 := 123456.56
	fmt.Println(numeroReal3)

	// String
	var str1 string = "Texto1"
	fmt.Println(str1)

	str2 := "Texto2"
	fmt.Println(str2)

	// com aspas simples, retorna o código do caracter na tabela ascii
	char := 'B'
	fmt.Println(char)

	// Valor ZERO, declarada sem atribuição de valor
	var Texto string
	fmt.Println(Texto)

	var int int16
	fmt.Println(int)

	// Boolean
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)

	// ERROR
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
