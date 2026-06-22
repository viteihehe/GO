package main

import (
	"fmt"
	"time"
)

func Filho(canal chan int) {
	for i := 1; ; i++ {
		canal <- i
	}
}

func Pai(canal chan int) {
	for {
		valor := <-canal
		fmt.Printf("Estou recebendo a mensande %d de meu filho\n", valor)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var canal chan int = make(chan int)

	go Filho(canal)
	go Pai(canal)

	var temp string
	fmt.Scan(&temp)

}
