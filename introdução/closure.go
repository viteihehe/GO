package main

import "fmt"

//Declaração basica de função

func decre2 (b int32) int32 {
	b -= 2
	return b
}

func main()	{
	//clossure é a capacidade de criar e chamar funções dentro de outras

	var x int32 = 0

	incre := func(a int32) int32 {
		a++
		return a
	}
	//defer é a capacidade de escalonar a função ou seja
	//adiar ela, pois o escalonamento funciona ao contrario
	//Ex:  x = 0, 
	// defer decre2(x) Roda dps devudo o defer
	// incre(x) Roda primeiro
	fmt.Println("X antes: ",x," X dps: ", incre(x)," X decre: ", decre2(x));
}