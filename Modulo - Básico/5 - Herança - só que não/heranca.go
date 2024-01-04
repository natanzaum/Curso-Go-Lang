package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Silva", 20}

	e1 := estudante{p1, "ADS", "UFMG"}

	fmt.Println(e1)
	fmt.Println(e1.sobrenome)
	fmt.Println(e1.pessoa.nome)

}
