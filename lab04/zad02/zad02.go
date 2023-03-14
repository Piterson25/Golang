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
	tab1 := [][]float64{{1, 2}, {3, 1}}
	tab2 := [][]float64{{0, 3}, {2, 1}}

	printMatrix(tab1)
	printMatrix(tab2)

	var result = [][]float64{}

	for i := 0; i < len(tab1); i++ {
		var row = []float64{}
		for j := 0; j < len(tab2[i]); j++ {
			row = append(row, tab1[i][j]*tab2[i][j])
		}

		result = append(result, row)
	}

	printMatrix(result)

}
