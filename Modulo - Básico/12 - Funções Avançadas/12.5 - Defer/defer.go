package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func main() {

	defer fmt.Println("Atrasa tudo, até print")

	defer funcao1() //Defer adia a execução de uma função até o ultimo momento possível
	funcao2()

}
