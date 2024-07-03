package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	fmt.Println("Price-Calculator-v6-json-tags-struct-in-struct")
	taxRates := []float64{0, 0.07, 0.01, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.Filemanager{
			InputFile:  "prices.txt",
			OutputFile: fmt.Sprintf("result_%.0f", taxRate*100),
		}
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		job.Process()
	}
}
