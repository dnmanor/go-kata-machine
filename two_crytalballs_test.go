package main

import (
	"math"
	"math/rand"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	// Original test case
	t.Run("Two crystal balls with 10000-element array", func(t *testing.T) {
		idx := int(math.Floor(rand.Float64() * 10000))
		data := make([]bool, 10000)
		for i := idx; i < 10000; i++ {
			data[i] = true
		}

		result := TwoCrystalBalls(data)
		if result != idx {
			t.Errorf("Expected result to be %d, but got %d", idx, result)
		}
	})

	// Test with an empty array
	t.Run("Two crystal balls with empty array", func(t *testing.T) {
		result := TwoCrystalBalls([]bool{})
		if result != -1 {
			t.Errorf("Expected result to be -1, but got %d", result)
		}
	})

	// Test with an array where all elements are true
	t.Run("Two crystal balls with all true array", func(t *testing.T) {
		data := make([]bool, 10000)
		for i := range data {
			data[i] = true
		}

		result := TwoCrystalBalls(data)
		if result != 0 {
			t.Errorf("Expected result to be 0, but got %d", result)
		}
	})

	// Test with an array where all elements are false
	t.Run("Two crystal balls with all false array", func(t *testing.T) {
		data := make([]bool, 10000)

		result := TwoCrystalBalls(data)
		if result != -1 {
			t.Errorf("Expected result to be -1, but got %d", result)
		}
	})

	// Add more test cases as needed
}
