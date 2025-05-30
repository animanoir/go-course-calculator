package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")
	prices := []string{} // Initialize the slice to store prices
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price) // Append price to the prices slice
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
