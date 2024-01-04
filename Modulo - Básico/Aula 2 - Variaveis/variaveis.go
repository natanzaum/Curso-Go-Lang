package main

import "fmt"

func main() {
	var variavel1 string = "Isso é uma string"
	fmt.Println(variavel1)

	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste3"
		variavel4 string = "teste4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "teste5", "dsds"

	fmt.Println(variavel5, variavel6)
}
