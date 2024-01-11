package main

import (
	"fmt"
	"time"
)

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}

func main() {
	canal := escrever("Ola Mundo!")

	for i := 0; i < 100; i++ {
		fmt.Println(<-canal)
	}
}
