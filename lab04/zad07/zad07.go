package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows        = 30
	columns     = 60
	generations = 100
)

func main() {
	board := generateBoard(rows, columns)
	for i := 0; i < generations; i++ {
		printBoard(board)
		board = updateBoard(board)
		time.Sleep(500 * time.Millisecond)
	}
}

func generateBoard(rows, columns int) [][]bool {
	board := make([][]bool, rows)
	for i := range board {
		board[i] = make([]bool, columns)
		for j := range board[i] {
			board[i][j] = rand.Intn(2) == 0
		}
	}
	return board
}

func printBoard(board [][]bool) {
	fmt.Print("\033[2J")
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				fmt.Printf("\033[%d;%dH#", i+1, j*2+1)
			} else {
				fmt.Printf("\033[%d;%dH ", i+1, j*2+1)
			}
		}
	}
	fmt.Printf("\033[%d;%dH", len(board)+1, 0)
}

func updateBoard(board [][]bool) [][]bool {
	newBoard := make([][]bool, len(board))
	for i := range newBoard {
		newBoard[i] = make([]bool, len(board[i]))
		for j := range newBoard[i] {
			aliveNeighbors := countAliveNeighbors(board, i, j)
			if board[i][j] {
				if aliveNeighbors == 2 || aliveNeighbors == 3 {
					newBoard[i][j] = true
				}
			} else {
				if aliveNeighbors == 3 {
					newBoard[i][j] = true
				}
			}
		}
	}
	return newBoard
}

func countAliveNeighbors(board [][]bool, row, col int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := row + i
			c := col + j
			if r < 0 || r >= len(board) || c < 0 || c >= len(board[row]) {
				continue
			}
			if board[r][c] {
				count++
			}
		}
	}
	return count
}
