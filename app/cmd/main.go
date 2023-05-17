package main

import (
	"ProductAnalyzerGo/app/internal/app"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path")
	}

	filePath := os.Args[1]

	err := app.App(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
