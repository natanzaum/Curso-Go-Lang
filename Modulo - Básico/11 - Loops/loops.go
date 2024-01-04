package main

import (
	"fmt"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i: ", i)
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j: ", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"Jose", "Mayara", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}
