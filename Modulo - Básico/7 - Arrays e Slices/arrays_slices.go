package main

import "fmt"

func main() {

	var array [5]int

	fmt.Println(array)

	array[1] = 10
	fmt.Println(array)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6} //Não deixa dinamico, só deixa com a quantidade especificada

	fmt.Println(array3)

	//slice, tipo array, mas mais flexivel

	slice := []int{10, 10, 12, 14, 16, 18, 20}
	fmt.Println(slice)

	slice = append(slice, 50)

	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)

}
