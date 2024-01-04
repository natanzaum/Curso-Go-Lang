package main

import "fmt"

type forma interface {
	area() float64
}

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func (c circulo) area() float64 {
	return 3.14 * (c.raio * c.raio)
}

func escreverArea(f forma) {
	fmt.Println("A area da forma é: ", f.area())
}

func main() {
	r := retangulo{10, 20}
	escreverArea(r)

	c := circulo{200}
	escreverArea(c)

}
