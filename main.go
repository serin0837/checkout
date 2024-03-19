package main

import (
	"checkout-kata/utils"
	"fmt"
)

// Checkout function scans the item and calculates the total price.
func Checkout(scanItems []string) int {

	// Init this week SKU
	sku := utils.ThisWeekSKU()

	for _, scanItem := range scanItems {
		utils.Scan(scanItem, sku)
	}

	totalPrice := utils.TotalPrice(sku)

	return totalPrice
}

func main() {
	checkoutPrice := Checkout([]string{"B", "A", "B"})
	fmt.Printf("Total price of scanned item is: %d", checkoutPrice)
}
