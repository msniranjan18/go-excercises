package main

import (
	"fmt"
)

func main() {
	fmt.Println("Price-Calculator")
	taxRates := []float64{0, 0.07, 0.01, 0.15}
	prices := []float64{10, 20, 30}

	result := make(map[float64][]float64)
	for _, taxRate := range taxRates {
		tempSlice := make([]float64, len(prices))
		for idx, price := range prices {
			tempSlice[idx] = price * (1 + taxRate)
		}
		result[taxRate] = tempSlice
	}

	fmt.Println("result", result)
}
