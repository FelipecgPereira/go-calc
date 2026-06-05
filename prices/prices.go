package prices

import (
	"fmt"

	"com.github/FelipecgPereira/go-calc/conversion"
	"com.github/FelipecgPereira/go-calc/iomanager"
)

type TaxIncludedPrice struct {
	InputPrices       []float64
	TaxRate           float64
	TaxIncludedPrices map[string]string
	IOManager         iomanager.IOManager
}

func (inputJob *TaxIncludedPrice) Process(doneChan chan bool, errorChan chan error) {
	err := inputJob.LoadData()
	if err != nil {
		errorChan <- err
		return
	}

	result := make(map[string]string)
	for _, price := range inputJob.InputPrices {
		taxIncludedPrice := price * (1 + inputJob.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	inputJob.TaxIncludedPrices = result

	inputJob.IOManager.Write(inputJob)
	doneChan <- true
}

func (inputJob *TaxIncludedPrice) LoadData() error {
	lines, err := inputJob.IOManager.Read()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	inputJob.InputPrices = prices

	return nil
}

func NewTaxIncludePriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		InputPrices:       []float64{},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]string),
		IOManager:         iom,
	}
}
