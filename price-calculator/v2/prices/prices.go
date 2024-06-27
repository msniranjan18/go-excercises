package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	Prices            []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		Prices:  []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	job.TaxIncludedPrices = result
	job.Display()
}

func (job *TaxIncludedPriceJob) Display() {
	fmt.Printf("Tax Rate: [%v]\n", job.TaxRate)
	fmt.Printf("Taxed Price List: %+v\n\n", job.TaxIncludedPrices)
}
