package csvParse

import (
	"encoding/csv"
	"os"
)

func Parse(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(csvFile)
	return reader.ReadAll()
}
