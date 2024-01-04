package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar")
	fmt.Println("Dados do usuario: Nome:", u.nome, "Idade:", u.idade)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario1", 20}
	usuario1.salvar()

	usuario1.fazerAniversario()

	usuario1.salvar()

}
