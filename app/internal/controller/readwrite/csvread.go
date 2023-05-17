package readwrite

import (
	"ProductAnalyzerGo/app/internal/controller"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReadCSV(filePath string) ([]controller.Product, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("IOError")
		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	products := make([]controller.Product, len(records)-1)
	for i, record := range records[1:] {
		product := controller.Product{
			Name: record[0],
			Cost: record[1],
		}

		rating, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		product.Rating = rating

		products[i] = product
	}

	return products, nil
}
