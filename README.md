# ProductAnalyzerGo

ProductAnalyzerGo - это проект, который анализирует данные о продуктах из CSV и JSON файлов.

## стэк

- Go 
- Docker 

## запуск

go run main.go ../../<путь_к_CSV/JSON_файлу>

или в Docker среде

docker run -v "$(pwd)/resources/product.csv:/resources/product.csv" -p 8080:8080 product-analyzer



