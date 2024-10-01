package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	salario  float32
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Struct")

	var u1 usuario
	fmt.Println(u1)
	u1.nome = "Emerson"
	u1.idade = 52
	fmt.Println(u1)
	var u2 usuario
	u2.nome = "Mara"
	u2.idade = 50
	u2.salario = 15000.00
	fmt.Println(u2)
	fmt.Println(u1, u2)

	enderecoU3 := endereco{"Rua qualquer", 99}
	u3 := usuario{"Lipe", 21, 10000, enderecoU3}
	fmt.Println(u3)

	u4 := usuario{nome: "Amanda", idade: 18}
	fmt.Println(u4)
}
