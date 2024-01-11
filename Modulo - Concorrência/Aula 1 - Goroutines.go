package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo") //Goroutine (Executa a função, mas não espera ela terminar para executar a proxima linha)
	escrever("Programa em GO")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
