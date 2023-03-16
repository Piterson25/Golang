package main

import "fmt"

func matricesAreEqual(matrix1 [][]int64, matrix2 [][]int64) bool {
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		return false
	}

	for i := range matrix1 {
		for j := range matrix1[i] {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}

	return true
}

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

func printMatrix(matrix [][]int64) {
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

func main() {
	matrix1 := [][]int64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	matrix2 := [][]int64{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1}}

	result1 := [][]int64{}
	result2 := [][]int64{}

	if len(matrix1[0]) != len(matrix2) || len(matrix2[0]) != len(matrix1) {
		fmt.Println("Nie można pomnożyć tych macierzy")
		return
	} else {
		result1 = multiplyMatrix(matrix1, matrix2)
		result2 = multiplyMatrix(matrix2, matrix1)
	}

	fmt.Println("Wynik mnożenia macierzy pierwszej:")
	printMatrix(result1)

	fmt.Println("Wynik mnożenia macierzy drugiej:")
	printMatrix(result2)

	if matricesAreEqual(result1, result2) {
		fmt.Println("Mnożenie tych macierzy jest przemienne")
	} else {
		fmt.Println("Mnożenie tych macierzy nie jest przemienne")
	}
}
