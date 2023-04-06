package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// Test cases
	vals := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	// Test for existing value
	if !BinarySearch(69, vals) {
		t.Errorf("Expected true, but got false")
	}

	// Test for non-existing value
	if BinarySearch(1336, vals) {
		t.Errorf("Expected false, but got true")
	}

	// Test for existing value at the end
	if !BinarySearch(69420, vals) {
		t.Errorf("Expected true, but got false")
	}

	// Test for non-existing value at the end
	if BinarySearch(69421, vals) {
		t.Errorf("Expected false, but got true")
	}

	// Test for existing value at the beginning
	if !BinarySearch(1, vals) {
		t.Errorf("Expected true, but got false")
	}

	// Test for non-existing value at the beginning
	if BinarySearch(0, vals) {
		t.Errorf("Expected false, but got true")
	}
}
