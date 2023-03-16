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

func reverseMatrix(matrix [][]float64) [][]float64 {
	result := make([][]float64, len(matrix))

	for i := range result {
		result[i] = make([]float64, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

func addMatrixToReverse(matrix [][]float64) [][]float64 {
	reversedMatrix := reverseMatrix(matrix)
	result := make([][]float64, len(matrix))

	for i := range result {
		result[i] = make([]float64, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[i][j] = matrix[i][j] + reversedMatrix[i][j]
		}
	}

	return result
}

func main() {
	tab := [][]float64{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	printMatrix(tab)

	result := addMatrixToReverse(tab)

	printMatrix(result)
}
