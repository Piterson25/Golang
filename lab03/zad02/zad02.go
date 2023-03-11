package main

import (
	"fmt"
	"io"
	"os"
)

type seq struct {
	seq_len int
	number  int
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

func findMaxLength(arr []seq) (int, int) {
	max := arr[0]
	for _, value := range arr {
		if value.seq_len > max.seq_len {
			max = value
		}
	}
	return max.seq_len, max.number
}

func main() {

	seqs := []seq{}

	const max_range int = 3000

	for i := 1000; i <= max_range; i++ {
		seqs = append(seqs, seq{seq_len: len(collatz(i, []int{})), number: i})
		if i%1000 == 0 && i > 1000 {
			max, number := findMaxLength(seqs)
			s := fmt.Sprintf("%d-%d Najdłuższy ciąg Collatza ma liczba: %d (%d elementów)\n", i-1000, i, number, max)
			io.WriteString(os.Stdout, s)
			seqs = []seq{}
		}
	}

}
