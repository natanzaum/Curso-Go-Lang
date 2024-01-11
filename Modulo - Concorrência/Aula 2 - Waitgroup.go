package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(2) //quantidade de Goroutines rodando ao mesmo tempo

	go func() {
		escrever("Olá mundo")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programa em GO")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
