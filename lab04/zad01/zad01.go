package main

import (
	"fmt"
)

func main() {
	tab1 := [20]float64{}
	tab2 := [20]float64{}

	for i := 0; i < len(tab1); i++ {
		tab1[i] = 2.0
		tab2[i] = 3.0
	}

	var result float64 = 0

	for i := 0; i < len(tab1); i++ {
		result += tab1[i] + tab2[i]
	}

	fmt.Println(result)

}
