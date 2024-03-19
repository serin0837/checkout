package main

import (
	"checkout-kata/models"
	"checkout-kata/utils"
	"testing"
)

func TestCheckout(t *testing.T) {

	todaySku := utils.ThisWeekSKU()
	tmrSku := utils.ThisWeekSKU()

	nextweekSku := utils.ThisWeekSKU()

	// Directly modify the item's price, but next time we want to set it differently.
	nextweekSku[2].UnitPrice = 100

	testCases := []struct {
		scanItems     []string
		sku           []models.Item
		expectedPrice int
	}{
		// Test 1
		{
			[]string{"B", "A", "B"},
			todaySku,
			95,
		},
		// Test2
		{
			[]string{"B", "A", "B", "B", "C", "D"},
			tmrSku,
			160,
		},
		// Test3
		{
			[]string{"B", "A", "B", "C", "C"},
			nextweekSku,
			295,
		},
		// for negative testing
		// {
		// 	[]string{"B", "A", "B", "F", "J"},
		// 	nextweekSku,
		// 	295,
		// },
	}

	for _, tt := range testCases {
		got := Checkout(tt.scanItems, tt.sku)
		if got != tt.expectedPrice {
			t.Errorf("total price should calculate as expected. got: %v expected: %v", got, tt.expectedPrice)
		}
	}
}
