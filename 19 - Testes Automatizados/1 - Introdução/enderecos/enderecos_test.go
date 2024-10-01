package enderecos_test

import (
	"introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEnderecp(t *testing.T) {

	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Mário Covas", "Rodovia"},
		//{"Alameda Bem Fica", "Alameda"},
		//{"Praça 7 de Setembro", "Praça"},
		{"Estrada teste", "Estrada"},
		//{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := enderecos.TipoDeEndereco(cenario.retornoEsperado)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		}
	}
	// Dicas de uso go test
	// go test
	// go test ./...   								  	- testar todos o pacotes
	// go test -v       							    - cada função
	// go test --cover   				    			- Exibe % de cobertura
	// go test --coverprofile <name.txt>  - gera um arquivo com o resultado do teste
	// go tool cover --func=<name.txt>    - Exibe o percentual por pacote
	// go tool cover --html=<name.txt>    - Exibe de forma detalhada em html sobre onde não há cobertura de testes

}

func TestQualqu(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou...")
	}
}
