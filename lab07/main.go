package main

import (
	"flag"
	"fmt"
	"math"
)

type p struct {
	a, b, c float64
}

type T struct {
	x, y float64
}

func s(d *p) (T, bool) {
	if x := d.b*d.b - 4*d.a*d.c; x >= 0 && d.a != 0 {
		q := math.Sqrt(x)
		return T{(-d.b - q) / (2 * d.a), (-d.b + q) / (2 * d.a)}, true
	}
	return T{0, 0}, false
}

func main() {
	var a, b, c float64
	flag.Float64Var(&a, "a", 0, "Liczba a")
	flag.Float64Var(&b, "b", 0, "Liczba b")
	flag.Float64Var(&c, "c", 0, "Liczba c")
	flag.Parse()
	t, r := s(&p{a, b, c})
	fmt.Println(t, r)
}
