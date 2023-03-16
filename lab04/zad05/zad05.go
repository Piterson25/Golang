package main

import "fmt"

func multiplyMatrix(matrix1 [][]int64, matrix2 [][]int64) [][]int64 {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	cols2 := len(matrix2[0])

	result := make([][]int64, rows1)
	for i := range result {
		result[i] = make([]int64, cols2)
	}

	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

func main() {
	matrix1 := [][]int64{
		{1, 0, 2},
		{-1, 3, 1}}
	matrix2 := [][]int64{
		{3, 1},
		{2, 1},
		{1, 0}}

	result := [][]int64{}

	if len(matrix1[0]) != len(matrix2) {
		fmt.Println("Nie można pomnożyć tych macierzy")
		return
	} else {
		result = multiplyMatrix(matrix1, matrix2)
	}

	fmt.Println("Wynik mnożenia macierzy:")
	for i := range result {
		fmt.Println(result[i])
	}
}
