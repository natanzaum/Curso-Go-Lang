package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("Numero maior que 15")
	} else {
		fmt.Println("Numero menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero maior maior que 0")
	} else if numero < -10 {
		fmt.Println("Numero menor que -10")
	} else {
		fmt.Println("Numero entre 0 e -10")
	}

}
