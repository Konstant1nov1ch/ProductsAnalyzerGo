package readwrite

import (
	"encoding/json"
	"os"

	"ProductAnalyzerGo/app/internal/controller"
)

func ReadJSON(filePath string) ([]controller.Product, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []controller.Product
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
