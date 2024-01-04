package main

import "fmt"

func Somar(a int, b int) int {
	return a + b
}

func CalculosMatematicos(a, b int32) (int32, int32) {

	soma := a + b
	subtracao := a - b

	return soma, subtracao //Função com multiplos retornos
}

func main() {
	soma := Somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função dentro da função")
	}

	f()

	resultadoSoma, resultadoSubtracao := CalculosMatematicos(10, 5)

	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := CalculosMatematicos(10, 5) //Ignora o resultado do segundo Return
	fmt.Println(resultadoSoma2)
}
