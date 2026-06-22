package main

import (
	"fmt"
	"time"
)

func Ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func Pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func Mesa(c chan string) {
	for {
		texto := <-c
		fmt.Print(texto + "\n")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go Pong(c)
	go Mesa(c)
	go Ping(c)

	select {}
}
