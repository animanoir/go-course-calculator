package main

import (
	//"fmt"
	"go-course-calculator/prices"
	//"go-course-calculator/filemanager"
	"go-course-calculator/cmdmanager"
)

// The difference between an interface and a stuct is in their purpose:
// - struct is kind of like a cookie cutter — defines a concrete data type grouping data to create objects from it — defines what is something.
// - an interface establishes "rules" or "methods — defines what something can do, like a contract.
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}

//NOTE - Personal notes:
//NOTE - Arrays are like a fixed-size egg carton. I can decide upfront how many eggs it can hold and it will never change. Each slut is numbered and can only hold one egg.
//NOTE - Slices are like an expandable photo album. I can add more pages when needed. I can also create a "view" that shows certain pages of the album. Under the hood they are like "arrays" (references arrays), but provide flexibility to grow.
//NOTE - A map is like a dictionary or address book. I can look up things up by a key (like a name or ID) to find the corresponding value.
//NOTE DO NOT CREATE FILES FROM THE TERMINAL BECAUSE FUCKING WINDOWS CHANGES THE CHARACTER ENCODING
