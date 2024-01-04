package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Aula de pacotes")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("natan@gmail.com")
	fmt.Println(erro)

	erro = checkmail.ValidateFormat("natangmail.com")
	fmt.Println(erro)

}
