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
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)

	if result != expected {
		t.Errorf("Reverse(%q) == %q, expected %q", input, result, expected)
	}
}

// Benchmark for SwapCase function
func BenchmarkSwapCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwapCase("Hello, World")
	}
}

//Benchmark for Reverse function
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("Hello, World")
	}
}