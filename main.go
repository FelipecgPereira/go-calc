package main

import (
	"fmt"

	"com.github/FelipecgPereira/go-calc/filemanager"
	"com.github/FelipecgPereira/go-calc/prices"
)

func main() {
	taxRates := []float64{0.07, 0.08, 0.06, 0.05}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("tax_include_prices_%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludePriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

	}

	for index, _ := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Printf("Error processing price job: %v\n", err)
			}
		case <-doneChans[index]:
			fmt.Printf("Price job with tax rate %.0f%% completed.\n", taxRates[index]*100)
		}
	}

}
