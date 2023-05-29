package forest

import (
	"fmt"
	"math/rand"
)

func Simulation(forestWidth int, forestHeight int) (int, int) {
	var forest [][]int

	for i := 0; i < forestWidth; i++ {
		var row []int
		for j := 0; j < forestHeight; j++ {
			row = append(row, 0)
		}
		forest = append(forest, row)
	}

	trees := 0

	for x, v := range forest {
		for y := range v {
			treeOrNot := rand.Intn(2)
			forest[x][y] = treeOrNot

			if treeOrNot == 1 {
				trees += 1
			}

		}
	}

	updatedForest := fire(forest, rand.Intn(len(forest)), rand.Intn(len(forest[0])))

	burned := 0
	for x, v := range updatedForest {
		for y := range v {
			if updatedForest[x][y] == 1 {
				fmt.Print("\x1b[32m")
			} else if updatedForest[x][y] == 2 {
				burned += 1
				fmt.Print("\x1b[31m")
			} else {
				fmt.Print("\x1b[0m")
			}
			fmt.Print(updatedForest[x][y], " ")
		}
		fmt.Println()
	}
	fmt.Print("\x1b[0m")

	return trees, burned
}

func fire(forest [][]int, lightningX int, lightningY int) [][]int {
	if lightningX >= 0 && lightningY >= 0 && lightningX < len(forest) && lightningY < len(forest[0]) {
		if forest[lightningX][lightningY] == 1 {
			forest[lightningX][lightningY] = 2
			fire(forest, lightningX-1, lightningY)
			fire(forest, lightningX, lightningY-1)
			fire(forest, lightningX+1, lightningY)
			fire(forest, lightningX, lightningY+1)
		}
	}

	return forest
}
