package main

import (
	"checkout-kata/models"
	"fmt"
)

// func NewItem(name string, unitPrice, specialPrice int) models.Item {
// 	return models.Item{
// 		Name:         name,
// 		UnitPrice:    unitPrice,
// 		SpecialPrice: specialPrice,
// 	}
// }

func ThisWeekSKU() []models.Item {
	return []models.Item{
		{
			Name:         "A",
			UnitPrice:    50,
			SpecialPrice: 130,
			Number:       0,
		},
		{
			Name:         "B",
			UnitPrice:    30,
			SpecialPrice: 45,
		},
		{
			Name:         "C",
			UnitPrice:    20,
			SpecialPrice: 0,
		},
		{
			Name:         "D",
			UnitPrice:    15,
			SpecialPrice: 0,
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

func TotalPrice() string {
	return "Hello, world"
}

func main() {
	fmt.Println(TotalPrice())

	// Let's scan the items

}
