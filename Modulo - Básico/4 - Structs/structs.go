package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	fmt.Println("Aula de struct")

	var u usuario
	fmt.Println(u)

	u.nome = "Lucas"
	u.sobrenome = "Bonifacio"
	u.idade = 56

	fmt.Println(u)
}
