package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	var e1 estudante
	e1.nome = "Emerson"
	e1.sobrenome = "Barina"
	e1.idade = 52
	e1.altura = 180
	e1.curso = "Sistemas"
	e1.faculdade = "A vida"

	fmt.Println(e1)

	e2 := estudante{pessoa{"Mara", "Tosato", 50, 160}, "Psicanálise", "Freudiada"}
	fmt.Println(e2)

	p1 := pessoa{"Lipe", "Barina", 21, 176}
	fmt.Println(p1)

	e3 := estudante{p1, "Engennharia", "USP"}
	fmt.Println(e3)

}
