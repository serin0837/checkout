package main

import "testing"

func TestTotalPrice(t *testing.T) {
	got := TotalPrice()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
