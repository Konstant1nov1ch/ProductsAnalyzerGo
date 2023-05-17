package analyzer

import (
	"ProductAnalyzerGo/app/internal/controller"
	"strconv"
)

func FindMostExpensiveProduct(products []controller.Product) *controller.Product {
	if len(products) == 0 {
		return nil
	}

	mostExpensive := &products[0]
	maxCost, _ := strconv.ParseFloat(mostExpensive.Cost, 64)
	for i := 1; i < len(products); i++ {
		cost, _ := strconv.ParseFloat(products[i].Cost, 64)
		if cost > maxCost {
			mostExpensive = &products[i]
			maxCost = cost
		}
	}

	return mostExpensive
}

func FindHighestRatedProduct(products []controller.Product) *controller.Product {
	if len(products) == 0 {
		return nil
	}

	highestRated := &products[0]
	for i := 1; i < len(products); i++ {
		if products[i].Rating > highestRated.Rating {
			highestRated = &products[i]
		}
	}

	return highestRated
}
