package main

import "fmt"
import "math"

//Declaração da interface
type geometria interface {
	area() float32
}

type quadrado struct {
	lado float32
}

type circulo struct {
	raio float32
}

//Ambas as formas geometricas retornam utilizam os metodos da interface geometria
func (q *quadrado) area() float32 {
	return q.lado*q.lado
}

func (c *circulo) area() float32 {
	return math.Pi*(c.raio*c.raio)
}

//Pode se criar funções que chame a interface geometria e elas irão poder imprimir qualquer valor que use os metodos
//da interface X.
func resul (g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {
	q := quadrado{lado: 4}
	var c circulo
	c.raio = 3

	resul(&q)
	resul(&c)

}