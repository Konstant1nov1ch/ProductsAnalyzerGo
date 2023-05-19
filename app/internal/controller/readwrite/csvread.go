package readwrite

import (
	"ProductAnalyzerGo/app/internal/controller"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

func ReadCSV(filePath string, bufferSize int, numWorkers int) ([]controller.Product, error) {
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

	var products []controller.Product
	productsCh := make(chan []controller.Product, numWorkers)
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for chunk := range productsCh {
				for _, product := range chunk {
					products = append(products, product)
				}
			}
		}()
	}

	chunk := make([]controller.Product, 0, bufferSize)
	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				productsCh <- chunk
				break
			}
			return nil, err
		}

		product := controller.Product{
			Name: records[0],
			Cost: records[1],
		}

		rating, err := strconv.Atoi(records[2])
		if err != nil {
			rating = 0
			continue
		}
		product.Rating = rating

		chunk = append(chunk, product)

		if len(chunk) == bufferSize {
			productsCh <- chunk
			chunk = make([]controller.Product, 0, bufferSize)
		}
	}

	close(productsCh)
	wg.Wait()

	return products, nil
}
