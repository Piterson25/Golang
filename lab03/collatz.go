package main

import "fmt"

type seq struct {
	seq    int
	number int
}

func findMinAndMax(arr []seq) (int, int) {
	max := arr[0]
	for _, value := range arr {
		if value.seq > max.seq {
			max = value
		}
	}
	return max.seq, max.number
}

func collatz(x int, arr []int) []int {
	arr = append(arr, x)
	if x == 1 {
		return arr
	} else if x%2 == 0 {
		return collatz(x/2, arr)
	}
	return collatz(x*3+1, arr)
}

func main() {

	thousand := []seq{{seq: len(collatz(1000, []int{})), number: 1000}}

	for i := 1000; i <= 100000; i++ {

		if i%1000 == 0 {
			max, number := findMinAndMax(thousand)
			fmt.Println(i-1000, "-", i-1, "Maksymalna liczba:", number, "ma", max, "elementow")
			thousand = []seq{}
		} else {
			thousand = append(thousand, seq{seq: len(collatz(i, []int{})), number: i})
		}
	}

}
