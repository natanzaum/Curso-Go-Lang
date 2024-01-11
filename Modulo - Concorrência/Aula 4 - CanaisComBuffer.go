package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) //Canal com buffer. Só bloqueia quando atingir a capacidade máxima

	canal <- "Olá mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
