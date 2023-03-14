package main

import (
	"fmt"
)

func printMatrix(matrix [][]float64) {
	for i := 0; i < len(matrix); i++ {
		var row = []float64{}
		for j := 0; j < len(matrix[i]); j++ {
			row = append(row, matrix[i][j])
		}
		fmt.Println(row)
	}
	fmt.Println("")
}

func main() {
	tab1 := [][]float64{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8}}

	printMatrix(tab1)

}
