package main

import "fmt"

type pessoa struct {
	Nome string
	idade int
}

func main() {
	var x int32 = 1
	fmt.Println("Hello, World!", x, "Testando")
	fmt.Println("1"+"1 e ",x+x)
	var v [5] int8
	v[2] = 2
	fmt.Println(v)
	pessoa := pessoa{"Daniel", 18}
	fmt.Println(pessoa)
} 