package main

import "fmt"

func diaDaSemana(numero int) string {
	var diaDaSemana string
	switch numero {
	case 1:
		diaDaSemana = "Domingo"
	case 2:
		diaDaSemana = "Segunda-feira"
	case 3:
		diaDaSemana = "Terça-feira"
	default:
		diaDaSemana = "dia indefinido"
	}
	return diaDaSemana
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	default:
		diaDaSemana = "dia indefinido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Estruturas de Controle - Switch")
	dia := diaDaSemana(2)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
