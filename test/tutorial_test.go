package test

import "testing"

func calc(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	if calc(2, 3) != 5 {
		t.Error("Expected 2 + 3 equal 5")
	}
}
