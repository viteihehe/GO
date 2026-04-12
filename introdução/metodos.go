package main

import "fmt"

type Retangulo struct {
	largura, altura int32
}


//Struct/Objeto a qual ela pertence () nome () retorno
func (r *Retangulo) area() int32 {
	return r.altura*r.largura
}

func main() {
	var r Retangulo
	fmt.Println("Insira a largura:")
	fmt.Scanf("%d %d", &r.largura, &r.altura)

	area := r.area()

	fmt.Println("Area ",area)

	return
}