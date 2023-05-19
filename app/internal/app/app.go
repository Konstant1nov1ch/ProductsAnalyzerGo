package app

import (
	"ProductAnalyzerGo/app/internal/controller"
	"ProductAnalyzerGo/app/internal/controller/analyzer"
	"ProductAnalyzerGo/app/internal/controller/readwrite"
	"fmt"
	"strings"
)

func App(filePath string, numWorkers int) error {
	var products []controller.Product
	var err error
	// Размер буфера я решил обрабатывать файл чанками по 10 строк (на очень больших файлах можно увеличить этот параметр)
	var bufferSize = 10
	if strings.HasSuffix(filePath, ".json") {
		// Чтение данных из JSON
		products, err = readwrite.ReadJSON(filePath, bufferSize, numWorkers)
	} else if strings.HasSuffix(filePath, ".csv") {
		// Чтение данных из CSV
		products, err = readwrite.ReadCSV(filePath, bufferSize, numWorkers)
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
