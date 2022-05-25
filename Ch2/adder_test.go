package main

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 5

	if sum != expected {
		t.Errorf("Expected: %d, Got: %d", expected, sum)
	}
}
