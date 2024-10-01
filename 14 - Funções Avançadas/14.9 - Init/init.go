package main

import "fmt"

var n int

func init() {
	fmt.Println("Função INIT sendo executada...")
	n = 10
}

func main() {
	fmt.Println("Função MAIN sendo executada...")
	fmt.Println(n)
}
