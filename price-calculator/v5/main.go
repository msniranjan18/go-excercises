package main

import (
	"fmt"

	"example.com/price-calculator/prices"
)

func main() {
	fmt.Println("Price-Calculator-v2")
	taxRates := []float64{0, 0.07, 0.01, 0.15}

	for _, taxRate := range taxRates {
		job := prices.NewTaxIncludedPriceJob(taxRate)
		job.Process()
	}
}
