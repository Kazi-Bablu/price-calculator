package main

import (
	"fmt"

	cmdmanager "example.com/price-calculator/cmdManager"
	"example.com/price-calculator/prices"
)

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRate {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		pricesJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := pricesJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
