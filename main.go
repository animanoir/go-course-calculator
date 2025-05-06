package main

import (
	"go-course-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	// consoleReader := bufio.NewReader(os.stdin)

}

//NOTE - Personal notes:
//NOTE - Arrays are like a fixed-size egg carton. I can decide upfront how many eggs it can hold and it will never change. Each slut is numbered and can only hold one egg.
//NOTE - Slices are like an expandable photo album. I can add more pages when needed. I can also create a "view" that shows certain pages of the album. Under the hood they are like "arrays" (references arrays), but provide flexibility to grow.
//NOTE - A map is like a dictionary or address book. I can look up things up by a key (like a name or ID) to find the corresponding value.
