package main

import "fmt"

func funcao1() {
	fmt.Println("Executando 1...")
}
func funcao2() {
	fmt.Println("Executando 2...")
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer fmt.Println("Media calculada. Resultado será apresentado!")
	fmt.Println("Entrando na função que verifica se aluno está aprovado...")

	media := (n1 + n2) / 2
	return media >= 6
	// if media >= 6 {
	// 	return true
	// }
	// return false
}

func main() {
	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))
}
