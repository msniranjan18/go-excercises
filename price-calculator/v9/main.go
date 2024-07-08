package main

import (
	"fmt"
	"sync"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	fmt.Println("Price-Calculator-v8-Use-of-Go-Routines-and-waitgroup")
	taxRates := []float64{0, 0.07, 0.01, 0.15}
	wg := new(sync.WaitGroup)

	for _, taxRate := range taxRates {
		wg.Add(1)
		iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f", taxRate*100))
		//iom := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(iom, taxRate)
		go job.Process(wg)
	}
	wg.Wait()

}
