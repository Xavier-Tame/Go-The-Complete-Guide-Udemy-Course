package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	balanceStr := fmt.Sprintf("%.2f", value)
	os.WriteFile(filename, []byte(balanceStr), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000.00, errors.New("file not found") // default balance if file doesn't exist
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000.00, errors.New("failed to parse balance") // default balance if parsing fails
	}
	return value, nil
}
