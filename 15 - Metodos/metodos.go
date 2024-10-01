package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	// método sempre estão vinculado à um struct
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	// if u.idade >= 18 {
	// 	return true
	// } else {
	// 	return false
	// }
	return u.idade >= 18
}

func exibeMaiorOuMenorDeIdade(nome string, maior bool) {
	if maior {
		fmt.Printf("%s é maior de idade\n", nome)
	} else {
		fmt.Printf("%s é menor de idade\n", nome)
	}
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"User 1", 40}
	usuario1.salvar()
	exibeMaiorOuMenorDeIdade(usuario1.nome, usuario1.maiorDeIdade())

	usuario2 := usuario{"User 2", 17}
	usuario2.salvar()
	exibeMaiorOuMenorDeIdade(usuario2.nome, usuario2.maiorDeIdade())
	fmt.Println(usuario2.idade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
	exibeMaiorOuMenorDeIdade(usuario2.nome, usuario2.maiorDeIdade())

}
