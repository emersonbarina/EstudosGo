package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA !=  PARALELISMO
	go escrever("Olá mundo!") // goroutine
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
