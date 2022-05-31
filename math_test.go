package main

import "testing"

func TestSoma(t *testing.T) {
	result := Sum(15, 15)
	if result != 30 {
		t.Errorf("Sum result is invalid; Result: %d. Expected: %d", result, 30)
	}
}
