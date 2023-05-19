package readwrite

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"ProductAnalyzerGo/app/internal/controller"
)

func ReadJSON(filePath string, bufferSize int, numWorkers int) ([]controller.Product, error) {
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

	var products []controller.Product
	decoder := json.NewDecoder(file)

	var wg sync.WaitGroup
	chunkSize := bufferSize / numWorkers
	chunks := make(chan []controller.Product)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for buffer := range chunks {
				// Обрабатываем каждый продукт из буфера
				for _, product := range buffer {
					product.Rating = sanitizeRating(product.Rating)
					products = append(products, product)
				}
			}
		}()
	}

	for {
		// Читаем буфер продуктов размером bufferSize
		buffer := make([]controller.Product, 0, chunkSize)
		if err := decoder.Decode(&buffer); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		// Отправляем буфер в канал для обработки горутинами
		chunks <- buffer

		if len(buffer) < bufferSize {
			break
		}
	}

	close(chunks)
	wg.Wait()

	return products, nil
}

func sanitizeRating(rating int) int {
	if rating < 0 {
		return 0
	}
	return rating
}
