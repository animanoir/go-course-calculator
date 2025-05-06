package main

import (
	// "bufio"
	"fmt"
)

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	result := make(map[float64][]float64) // make() creates an empty map
	for _, taxRate := range taxRates {
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}
	// consoleReader := bufio.NewReader(os.stdin)
	fmt.Print(result)

}

//NOTE - Personal notes:
//NOTE - Arrays are like a fixed-size egg carton. I can decide upfront how many eggs it can hold and it will never change. Each slut is numbered and can only hold one egg.
//NOTE - Slices are like an expandable photo album. I can add more pages when needed. I can also create a "view" that shows certain pages of the album. Under the hood they are like "arrays" (references arrays), but provide flexibility to grow.
//NOTE - A map is like a dictionary or address book. I can look up things up by a key (like a name or ID) to find the corresponding value.
