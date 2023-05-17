package main

import (
	"ProductAnalyzerGo/app/internal/app"
	"log"
	"os"
)

func main() {
	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
		log.Fatal("Please provide a file path")
	}

	err := app.App(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
