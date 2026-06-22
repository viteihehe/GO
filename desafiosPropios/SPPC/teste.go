package main

import (
	"fmt"
)

func main() {
	var valor int
	fmt.Print("Insira o valor:\n")
	fmt.Scan(&valor)

	fmt.Printf("Valor em decimal: %d\n", valor)
	valor = ConverterBin(valor)
	fmt.Printf("Valor em decimal: %d\n", valor)

	return
}

func ConverterBin(valor int) int {
	var valorBin int = 0
	decimal := 1
	var i int = valor
	for i >= 1 {
		valorBin += (i % 2) * decimal
		decimal *= 10
		i /= 2
	}
	fmt.Printf("%d\n", valorBin)
	return valorBin
}
