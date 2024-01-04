package main

import "fmt"

func AlunoAprovAado(n1, n2 float64) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media é exatamente 6!!!") //Para a exevução do programa
}

func recuperarExecucao() {
	if r := recover(); r != nil { //Recupera a execução do programa
		fmt.Println("Recuperou a execução")
	}

}

func main() {
	fmt.Println(AlunoAprovAado(6, 6))
}
