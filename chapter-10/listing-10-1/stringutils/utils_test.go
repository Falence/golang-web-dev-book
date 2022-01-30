package stringutils

import (
	"testing"
)

// Test case for the SwapCase function
func TestSwapCase(t *testing.T) {
	input, expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)

	if result != expected {
		t.Errorf("SwapCase(%q) == %q, expected %q", input, result, expected)
	}
}

// Test case for the Reverse function
func TestReverse(t *testing.T) {
	input, expected := "Hello, World", "dlroW olleH"
	result := Reverse(input)

	if result != expected {
		t.Errorf("Reverse(%q) == %q, expected %q", input, result, expected)
	}
}