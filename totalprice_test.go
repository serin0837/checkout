package main

import (
	"checkout-kata/models"
	"fmt"
	"testing"
)

func TestThisWeekSKU(t *testing.T) {
	got := ThisWeekSKU()
	want := []models.Item{
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

	// Check if slices have same length
	if len(got) != len(want) {
		t.Errorf("slices are not equal length")
	}

	for i, item := range got {
		if item.Name != want[i].Name {
			t.Errorf("slices are not equal. got: %q want: %q", item.Name, want[i].Name)
		}
	}
}

func TestScan(t *testing.T) {

	itemAName := "A"

	// InitSKU
	sku := ThisWeekSKU()
	got := Scan(itemAName, sku)

	want := []models.Item{
		{
			Name:         "A",
			UnitPrice:    50,
			SpecialPrice: 130,
			Number:       1,
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

	fmt.Printf("got: %+v\n", got)
	fmt.Printf("want: %+v\n", want)

	if got[0].Number != want[0].Number {
		t.Errorf("number of item should be same. got: %v want: %v", got[0].Number, want[0].Number)
	}
	// TODO: Add more test case

	// Scan B twice
	Scan("B", sku)
	got2 := Scan("B", sku)

	if got2[1].Number != 2 {
		t.Errorf("number of item should be same as expected number. got: %v want: %v", got[1].Number, 2)
	}
}

func TestTotalPrice(t *testing.T) {
	testSKU := []models.Item{
		{
			Name:         "A",
			UnitPrice:    50,
			SpecialPrice: 130,
			Number:       1,
		},
		{
			Name:         "B",
			UnitPrice:    30,
			SpecialPrice: 45,
			Number:       1,
		},
		{
			Name:         "C",
			UnitPrice:    20,
			SpecialPrice: 0,
			Number:       1,
		},
		{
			Name:         "D",
			UnitPrice:    15,
			SpecialPrice: 0,
			Number:       1,
		},
	}

	got := TotalPrice(testSKU)
	expected := 115
	if got != expected {
		t.Errorf("total price should calculate as expected. got: %v expected: %v", got, expected)
	}

	testSKU2 := []models.Item{
		{
			Name:         "A",
			UnitPrice:    50,
			SpecialPrice: 130,
			Number:       15,
		},
		{
			Name:         "B",
			UnitPrice:    30,
			SpecialPrice: 45,
			Number:       20,
		},
		{
			Name:         "C",
			UnitPrice:    20,
			SpecialPrice: 0,
			Number:       30,
		},
		{
			Name:         "D",
			UnitPrice:    15,
			SpecialPrice: 0,
			Number:       60,
		},
	}

	got2 := TotalPrice(testSKU2)
	expected2 := 2850
	if got2 != expected2 {
		t.Errorf("total price should calculate as expected here. got: %v expected: %v", got2, expected2)
	}

	// Test case for discount price
	testSKU3 := []models.Item{
		{
			Name:          "A",
			UnitPrice:     50,
			SpecialPrice:  130,
			SpecialNumber: 3,
			Number:        3,
		},
		{
			Name:          "B",
			UnitPrice:     30,
			SpecialPrice:  45,
			SpecialNumber: 2,
			Number:        2,
		},
		{
			Name:         "C",
			UnitPrice:    20,
			SpecialPrice: 0,
			Number:       1,
		},
		{
			Name:         "D",
			UnitPrice:    15,
			SpecialPrice: 0,
			Number:       1,
		},
	}

	got3 := TotalPrice(testSKU3)
	expected3 := 210
	if got3 != expected3 {
		t.Errorf("total price should calculate as expected. got: %v expected: %v", got3, expected3)
	}

	testSKU4 := []models.Item{
		{
			Name:          "A",
			UnitPrice:     50,
			SpecialPrice:  130,
			SpecialNumber: 3,
			Number:        15,
		},
		{
			Name:          "B",
			UnitPrice:     30,
			SpecialPrice:  45,
			SpecialNumber: 2,
			Number:        20,
		},
		{
			Name:          "C",
			UnitPrice:     20,
			SpecialPrice:  0,
			SpecialNumber: 0,
			Number:        1,
		},
		{
			Name:          "D",
			UnitPrice:     15,
			SpecialPrice:  0,
			SpecialNumber: 0,
			Number:        1,
		},
	}

	got4 := TotalPrice(testSKU4)
	expected4 := 1135
	if got4 != expected4 {
		t.Errorf("total price should calculate as expected. got: %v expected: %v", got4, expected4)
	}

}
