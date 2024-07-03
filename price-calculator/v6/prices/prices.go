package prices

import (
	"fmt"
	"strconv"

	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.Filemanager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	Prices            []float64               `json:"price"`
	TaxIncludedPrices map[string]string       `json:"price_including_tax"`
}

func NewTaxIncludedPriceJob(fm filemanager.Filemanager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPricesFromFile()
	result := make(map[string]string)
	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	err := job.IOManager.WriteJson(job)
	if err != nil {
		fmt.Println(err)
	}
	job.Display()
}

func (job *TaxIncludedPriceJob) Display() {
	fmt.Printf("Tax Rate: [%v]\n", job.TaxRate)
	fmt.Printf("Price List: %+v\n", job.Prices)
	fmt.Printf("Taxed Price List: %+v\n\n", job.TaxIncludedPrices)
}

func (job *TaxIncludedPriceJob) LoadPricesFromFile() {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, line := range lines {
		v, _ := strconv.ParseFloat(line, 64)
		job.Prices = append(job.Prices, float64(v))
	}
}
