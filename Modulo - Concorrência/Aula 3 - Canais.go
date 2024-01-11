package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	for {
		mensagem, aberto := <-canal //Recebe a mensagem do canal e continua a execução do programa. Verifica se ele esta aberto
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	/* Outra forma de fazer sem usar o "if !aberto"
	for mensagem := range canal{
		fmt.Println(mensagem)
	}
	*/
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //Envia a mensagem pro canal
		time.Sleep(time.Second)
	}

	close(canal)
}

//deadlock => Programa esta esperando algo no canal, mas que nunca vai chegar. Ai trava a execução
