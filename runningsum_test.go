package main

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		//from leetcode
	}

	for _, test := range tests {
		result := runningSum(test.input)

		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.output, result)
		}
	}
}
