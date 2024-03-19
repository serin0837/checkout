package utils

import (
	"checkout-kata/models"
	"log"
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
	found := false

	for i := range sku {
		// If existing item and scanned name is same, increase number
		if sku[i].Name == itemName {
			sku[i].Number++
			found = true
		}
	}
	if !found {
		log.Fatal("item name is not found")
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
		// if dividable, get quotient and remainder
		// then calculate total discount price.
		if skuItem.SpecialNumber != 0 {
			divided, remainder := divideAndGetRemainder(skuItem.Number, skuItem.SpecialNumber)
			discountedPrice = discountedPrice + divided*skuItem.SpecialPrice

			// change sku item number to remainder to calculate left over item.
			skuItem.Number = remainder

		}
		totalPrice = totalPrice + skuItem.UnitPrice*skuItem.Number + discountedPrice

	}
	return totalPrice
}

func divideAndGetRemainder(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}
