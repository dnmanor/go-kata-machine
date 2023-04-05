package main

import "testing"

func TestLinearSearch(t *testing.T) {
	// Test case 1: needle value in the haystack.
	needle := 5
	haystack := []int{1, 3, 5, 7, 9}
	if !LinearSearch(needle, haystack) {
		t.Errorf("LinearSearch(%d, %v) = false, expected true", needle, haystack)
	}

	// Test case 2: needle value not in the haystack.
	needle = 4
	haystack = []int{1, 3, 5, 7, 9}
	if LinearSearch(needle, haystack) {
		t.Errorf("LinearSearch(%d, %v) = true, expected false", needle, haystack)
	}

	// Test case 3: empty haystack.
	needle = 1
	haystack = []int{}
	if LinearSearch(needle, haystack) {
		t.Errorf("LinearSearch(%d, %v) = true, expected false", needle, haystack)
	}

	// Test case 4: haystack with one matching element.
	needle = 3
	haystack = []int{3}
	if !LinearSearch(needle, haystack) {
		t.Errorf("LinearSearch(%d, %v) = false, expected true", needle, haystack)
	}

	// Test case 5: haystack with one non-matching element.
	needle = 3
	haystack = []int{4}
	if LinearSearch(needle, haystack) {
		t.Errorf("LinearSearch(%d, %v) = true, expected false", needle, haystack)
	}

	// Test case 6: haystack with multiple matching elements.
	needle = 3
	haystack = []int{3, 3, 3}
	if !LinearSearch(needle, haystack) {
		t.Errorf("LinearSearch(%d, %v) = false, expected true", needle, haystack)
	}

	// Test case 7: haystack with multiple non-matching elements.
	needle = 2
	haystack = []int{3, 4, 5}
	if LinearSearch(needle, haystack) {
		t.Errorf("LinearSearch(%d, %v) = true, expected false", needle, haystack)
	}
}
