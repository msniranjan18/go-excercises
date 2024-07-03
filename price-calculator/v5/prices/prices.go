package prices

import (
	"fmt"
	"strconv"

	"example.com/price-calculator/filemanager"
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
	fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0.f", job.TaxRate*100))
	job.LoadPricesFromFile(fm)
	result := make(map[string]string)
	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	err := fm.WriteJson(job)
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

func (job *TaxIncludedPriceJob) LoadPricesFromFile(fm filemanager.Filemanager) {

	lines, err := fm.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, line := range lines {
		v, _ := strconv.ParseFloat(line, 64)
		job.Prices = append(job.Prices, float64(v))
	}
}
