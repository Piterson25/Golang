package forest

import (
	"encoding/csv"
	"fmt"
	"os"
)

func SaveResultsToCSV(filename string, results [][]int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range results {
		strRow := make([]string, len(row))
		for i, val := range row {
			strRow[i] = fmt.Sprintf("%d", val)
		}
		err := writer.Write(strRow)
		if err != nil {
			return err
		}
	}

	return nil
}
