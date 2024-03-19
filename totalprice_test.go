package main

import (
	"checkout-kata/models"
	"testing"
)

func TestTotalPrice(t *testing.T) {
	got := TotalPrice()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

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
