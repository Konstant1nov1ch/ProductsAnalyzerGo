# ProductAnalyzerGo

ProductAnalyzerGo - это проект, который анализирует данные о продуктах из CSV и JSON файлов.

## стэк

- Go 
- Docker 

## запуск

go run main.go ../../<путь_к_CSV/JSON_файлу>

или в Docker среде c переменной окружения FILE_PATH

 docker run -v "$(pwd)/resources:/app/resources" -p 8080:8080 --env FILE_PATH=/app/resources/product.json product-analyzer

## пример

![image](https://github.com/Konstant1nov1ch/ProductsAnalyzerGo/assets/105445251/b29c45b0-2ed1-4cbe-9533-5089de5875f9)


