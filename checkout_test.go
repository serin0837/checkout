package main

import (
	"testing"
)

func TestCheckout(t *testing.T) {
	got := Checkout([]string{"B", "A", "B"})
	expected := 95

	if got != expected {
		t.Errorf("checkout func should calculate as expected. got: %v expected: %v", got, expected)
	}
}
