package main

import (
	"fmt"

	"com.github/FelipecgPereira/go-calc/filemanager"
	"com.github/FelipecgPereira/go-calc/prices"
)

func main() {
	taxRates := []float64{0.07, 0.08, 0.06, 0.05}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("tax_include_prices_%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludePriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Error processing price job:", err)
		}
	}
}
