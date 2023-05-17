# ProductAnalyzerGo

ProductAnalyzerGo - это проект, который анализирует данные о продуктах из CSV и JSON файлов.

## стэк

- Go 
- Docker 

## запуск

go run main.go ../../<путь_к_CSV/JSON_файлу>

или в Docker среде

docker run -v "$(pwd)/resources/product.csv:/resources/product.csv" -p 8080:8080 product-analyzer

## пример

![image](https://github.com/Konstant1nov1ch/ProductsAnalyzerGo/assets/105445251/1833dc92-6cba-4d03-a180-d847d8d5646a)


