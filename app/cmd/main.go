package main

import (
	a "ProductAnalyzerGo/app/internal/app"
	"log"
)

func main() {
	//filePath := os.Getenv("FILE_PATH")
	//if filePath == "" {
	//	log.Fatal("Please provide a file path")
	//}
	//для быстрого теста
	filePath := "resources/productTest.csv"
	//количество потоков
	var numWorkers = 4
	err := a.App(filePath, numWorkers)
	if err != nil {
		log.Fatal(err)
	}
}
