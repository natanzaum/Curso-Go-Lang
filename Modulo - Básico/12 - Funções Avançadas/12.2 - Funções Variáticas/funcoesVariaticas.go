package main

import "fmt"

//Recebe de 1 a n numeros
//Só pode ter um variático por função e ele tem q ser o ultimo, caso tenha mais de um parametro
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total = total + numero
	}

	return total

}

func main() {

	soma := soma(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println(soma)
}
