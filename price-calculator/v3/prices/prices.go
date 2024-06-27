package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPricesFromFile("prices.txt")
	result := make(map[string]string)
	for _, price := range job.Prices {

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	job.Display()
}

func (job *TaxIncludedPriceJob) Display() {
	fmt.Printf("Tax Rate: [%v]\n", job.TaxRate)
	fmt.Printf("Price List: %+v\n", job.Prices)
	fmt.Printf("Taxed Price List: %+v\n\n", job.TaxIncludedPrices)
}

func (job *TaxIncludedPriceJob) LoadPricesFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, _ := strconv.ParseFloat(scanner.Text(), 64)
		job.Prices = append(job.Prices, float64(v))
	}

	if scanner.Err() != nil {
		fmt.Println(err)
		file.Close()
		return
	}
}
