package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Jr",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"])
}
