package prices

import (
	"fmt"
	"go-course-calculator/iomanager"
	"strconv"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`        // This will tell the JSON package that this key and its value should be ignored.
	TaxRate           float64             `json:"tax_rate"` // Adding this tags will tell the JSON package to writhe them instead of the real key.
	InputPrices       []float64           `json:"input:prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included"`
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string) // make() creates an empty map
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}
	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed.")
			return err
		}
		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices

	return err
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
