package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	fmt.Println("Price-Calculator-v8-Use-of-Go-Routines-and-channels-and-select")
	taxRates := []float64{0, 0.07, 0.01, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errChans[i] = make(chan error)
		iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f", taxRate*100))
		//iom := cmdmanager.New()
		job := prices.NewTaxIncludedPriceJob(iom, taxRate)
		go job.Process(doneChans[i], errChans[i])
	}
	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			fmt.Println(err)
		case <-doneChans[index]:
			fmt.Println("Done!!")
		}
	}

}
