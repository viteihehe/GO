package main

import "fmt"

const K int32 = 273

func main() {
	var tempk int32 = 500
	var c int32 = (tempk-K)
	fmt.Printf("Kelvin: %d Celsius: %d\n",tempk, c)
}