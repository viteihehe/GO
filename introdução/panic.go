//Panic: Utilizado para indicar uma mensagem de erro em caso de problema
//Recover: Resolver os panics

package main

import "fmt"

func main() {
	// 1. O defer precisa vir ANTES do erro
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Erro capturado:", x)
		}
	}()

	b := 0
	// if b == 0 {
	// 	// 2. Disparamos o panic manualmente ou o Go dispara sozinho no erro
	// 	panic("Divisão imprópria: tentativa de dividir por zero")
	// }
		
	fmt.Println(2 / b)
}
