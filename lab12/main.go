package main

import (
	"fmt"
	"lab12/forest"
)

func main() {
	numSimulations := 100
	forestWidth := 5
	forestHeight := 5

	var results [][]int

	for i := 0; i < numSimulations; i++ {
		trees, burned := forest.Simulation(forestWidth, forestHeight)
		fmt.Printf("Simulation %d - Remaining trees: %d | Burned: %d\n\n", i+1, trees, burned)

		result := []int{trees, burned}
		results = append(results, result)
	}

	filename := "forestSimulation.csv"
	err := forest.SaveResultsToCSV(filename, results)
	if err != nil {
		fmt.Println("Error saving results to CSV:", err)
	} else {
		fmt.Println("Successfully saved to CSV file!")
	}
}
