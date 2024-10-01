package main

import "fmt"

func recuperarExecucao() {
	//fmt.Println("Tentativa de recuperar execução...")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é 6")
}

func main() {
	fmt.Println("Executando Panic e Recover...")
	fmt.Println(alunoEstaAprovado(10, 2))
	fmt.Println("Após a média")
}
