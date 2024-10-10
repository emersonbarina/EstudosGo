package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função F")
	}

	f()

	var fp = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	fp("TESTE")

	resultado := fp("Teste resultado")
	fmt.Println(resultado)

	calc1, calc2 := calculosMatematicos(30, 20)
	fmt.Println(calc1, calc2)

	calc3, _ := calculosMatematicos(30, 20)
	fmt.Println(calc3)

}
