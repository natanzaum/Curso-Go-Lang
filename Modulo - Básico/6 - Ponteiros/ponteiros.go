package main

import "fmt"

func main() {
	var variavel int = 10
	var variavel2 = variavel

	fmt.Println(variavel, variavel2)

	variavel2++

	fmt.Println(variavel, variavel2)

	//Ponteiro = referência de memória

	var variavel3 int = 100
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro, &ponteiro, ponteiro, &variavel3)

}
