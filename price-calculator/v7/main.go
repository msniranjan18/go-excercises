package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {
	fmt.Println("Price-Calculator-v7-Use-of-Interface")
	taxRates := []float64{0, 0.07, 0.01, 0.15}

	for _, taxRate := range taxRates {
		//iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f", taxRate*100))
		iom := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(iom, taxRate)
		job.Process()
	}
}
