package main

import (
	"fmt"
	"math"
)

func delta(a float64, b float64, c float64) float64 {
	d := b*b - 4*a*c
	return d
}

func main() {
	fmt.Println("Podaj a")
	var a float64
	fmt.Scanln(&a)

	fmt.Println("Podaj b")
	var b float64
	fmt.Scanln(&b)

	fmt.Println("Podaj c")
	var c float64
	fmt.Scanln(&c)

	d := delta(a, b, c)

	if d < 0 {
		fmt.Println("Delta wyszła ujemna!")
	} else if d > 0 {
		x1 := -b - math.Sqrt(d)/2*a
		x2 := -b + math.Sqrt(d)/2*a
		fmt.Println("Dwa pierwiastki:", x1, x2)
	} else {
		x0 := -b / 2 * a
		fmt.Println("Podwójny pierwiastek:", x0)
	}

}
