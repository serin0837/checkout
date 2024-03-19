package main

import (
	"checkout-kata/utils"
	"fmt"
)

// Checkout function scans the item and calculates the total price.
func Checkout() int {

	// Init this week SKU
	sku := utils.ThisWeekSKU()

	// Scan B, A, B
	utils.Scan("B", sku)
	utils.Scan("A", sku)
	utils.Scan("B", sku)

	totalPrice := utils.TotalPrice(sku)

	return totalPrice
}

func main() {
	checkoutPrice := Checkout()
	fmt.Printf("Total price of scanned item is: %d", checkoutPrice)
}
