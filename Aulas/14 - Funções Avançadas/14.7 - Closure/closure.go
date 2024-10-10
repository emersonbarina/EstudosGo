package main

import "fmt"

func closure() func() {
	texto := "DENTRO DA FUNC CLOSURE"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "DENTRO DA MAIN"
	fmt.Println(texto)
	funcNova := closure()
	funcNova()
	fmt.Println(texto)

}
