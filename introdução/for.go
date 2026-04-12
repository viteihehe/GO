package main

import "fmt"

func main() {
	var i int32 = 0
	for i < 10 {
		break
		fmt.Println(i)
		i++
	}
	for y := 1; y <= 10; y++ {
		break
	}

	switch i {
	case 1:
	
	}

	v := [10]int32{1, 2, 3, 4, 5, 6,7, 8,9, 10}
	fatia := v[0:4]
	fmt.Println(fatia)
	fatia = append(fatia, 11, 12)
	fmt.Println(fatia)
	temp := make([]int32, 1)
	copy(temp, fatia)
	fmt.Println(temp)

	x := make(map[string]int)
	x["b"] = 20
	fmt.Println(x["a"])




}