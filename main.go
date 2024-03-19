package main

import (
	"checkout-kata/models"
	"checkout-kata/utils"
	"fmt"
)

// Checkout function scans the item and calculates the total price.
func Checkout(scanItems []string, sku []models.Item) int {

	for _, scanItem := range scanItems {
		utils.Scan(scanItem, sku)
	}

	totalPrice := utils.TotalPrice(sku)

	return totalPrice
}

func main() {
	// Init this week SKU
	sku := utils.ThisWeekSKU()

	checkoutPrice := Checkout([]string{"B", "A", "B"}, sku)

	fmt.Printf("Total price of scanned item is: %d", checkoutPrice)
}
