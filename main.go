package main

import (
	"checkout-kata/models"
	"checkout-kata/utils"
	"fmt"
)

func ThisWeekSKU() []models.Item {
	return []models.Item{
		{
			Name:      "A",
			UnitPrice: 50,
			// 3 for 130
			SpecialPrice:  130,
			SpecialNumber: 3,
			Number:        0,
		},
		{
			Name:      "B",
			UnitPrice: 30,
			// 2 for 45
			SpecialPrice:  45,
			SpecialNumber: 2,
			Number:        0,
		},
		{
			Name:         "C",
			UnitPrice:    20,
			SpecialPrice: 0,
			Number:       0,
		},
		{
			Name:         "D",
			UnitPrice:    15,
			SpecialPrice: 0,
			Number:       0,
		},
	}
}

// Scan function scan the item and keep track of numbers of item.
func Scan(itemName string, sku []models.Item) []models.Item {

	for i := range sku {
		// If existing item and scanned name is same, increase number
		if sku[i].Name == itemName {
			sku[i].Number++
		}

	}
	return sku
}

// TotalPrice calculates the total price of items in SKU.
func TotalPrice(sku []models.Item) int {
	var totalPrice int

	// Range each item in the sku
	for _, skuItem := range sku {
		var discountedPrice int

		// check if whole number is dividable by special number
		// check special price is not 0
		// if dividable, can we create sum of whole special price then minus the unitnumber
		if skuItem.SpecialNumber != 0 {
			divided, remainder := utils.DivideAndGetRemainder(skuItem.Number, skuItem.SpecialNumber)
			discountedPrice = discountedPrice + divided*skuItem.SpecialPrice
			fmt.Println(discountedPrice, "check discounted price")

			skuItem.Number = remainder

			fmt.Println("check remainder", remainder)
		}
		totalPrice = totalPrice + skuItem.UnitPrice*skuItem.Number + discountedPrice

	}
	return totalPrice
}

func main() {
	sku := ThisWeekSKU()

	// Scan B, A, B
	Scan("B", sku)
	Scan("A", sku)
	Scan("B", sku)

	totalPrice := TotalPrice(sku)

	fmt.Printf("Total price of scanned item is: %d", totalPrice)

}
