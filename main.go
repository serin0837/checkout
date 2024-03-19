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

func TotalPrice() string {
	return "Hello, world"
}

func main() {
	fmt.Println(TotalPrice())

	// Initialize this week stock
}
