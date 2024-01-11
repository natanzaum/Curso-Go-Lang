// Teste de unidade
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Otaviano Pena Forte", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Pedro Paulo", "Rodovia"},
		{"Praça Hebe Camargo", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.enderecoEsperado {
			t.Error("Tipo de endereço diferente do esperado!")
		}
	}
}
