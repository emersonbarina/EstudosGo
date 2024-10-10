package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco1 := enderecos.TipoDeEndereco("Rodovia do Imigrante")
	fmt.Println(tipoEndereco1)

	tipoEndereco2 := enderecos.TipoDeEndereco("Curva da Jurema")
	fmt.Println(tipoEndereco2)

}
