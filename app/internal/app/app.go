package app

import (
	"ProductAnalyzerGo/app/internal/controller"
	"ProductAnalyzerGo/app/internal/controller/analyzer"
	"ProductAnalyzerGo/app/internal/controller/readwrite"
	"fmt"
	"strings"
)

func App(filePath string) error {
	var products []controller.Product
	var err error

	if strings.HasSuffix(filePath, ".json") {
		// Чтение данных из JSON
		products, err = readwrite.ReadJSON(filePath)
	} else if strings.HasSuffix(filePath, ".csv") {
		// Чтение данных из CSV
		products, err = readwrite.ReadCSV(filePath)
	} else {
		return fmt.Errorf("unsupported file format: %s", filePath)
	}

	if err != nil {
		return err
	}

	// Проверка наличия продуктов
	if len(products) == 0 {
		fmt.Println("No products found in the file")
		return nil
	}

	// Поиск самого дорогого продукта
	mostExpensiveProduct := analyzer.FindMostExpensiveProduct(products)
	if mostExpensiveProduct != nil {
		fmt.Println("Most Expensive Product:", mostExpensiveProduct.Name)
	} else {
		fmt.Println("No products found in the file")
	}

	// Поиск продукта с самым высоким рейтингом
	highestRatedProduct := analyzer.FindHighestRatedProduct(products)
	if highestRatedProduct != nil {
		fmt.Println("Highest Rated Product:", highestRatedProduct.Name)
	} else {
		fmt.Println("No products found in the file")
	}

	return nil
}
