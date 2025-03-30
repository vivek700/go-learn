package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "sort empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "sort already sorted slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "sort reverse sorted slice",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "sort slice with duplicates",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "sort slice from main",
			input:    []int{20, 4, 80, 5, 3, 1, 6, 8, 2, 45, 89, 29, 19},
			expected: []int{1, 2, 3, 4, 5, 6, 8, 19, 20, 29, 45, 80, 89},
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Copy the input to avoid modifying the original
			input := make([]int, len(tc.input))
			copy(input, tc.input)

			// Call the function
			bubble_sort(input)

			// Check the result
			if !reflect.DeepEqual(input, tc.expected) {
				t.Errorf("bubble_sort(%v) = %v, expected %v", tc.input, input, tc.expected)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		slice    []int
		needle   int
		expected int
	}{
		{
			name:     "find in empty slice",
			slice:    []int{},
			needle:   5,
			expected: -1,
		},
		{
			name:     "find first element",
			slice:    []int{1, 2, 3, 4, 5},
			needle:   1,
			expected: 0,
		},
		{
			name:     "find last element",
			slice:    []int{1, 2, 3, 4, 5},
			needle:   5,
			expected: 4,
		},
		{
			name:     "find middle element",
			slice:    []int{1, 2, 3, 4, 5},
			needle:   3,
			expected: 2,
		},
		{
			name:     "element not present",
			slice:    []int{1, 2, 3, 4, 5},
			needle:   6,
			expected: -1,
		},
		{
			name:     "find in slice from main",
			slice:    []int{1, 2, 3, 4, 5, 6, 8, 19, 20, 29, 45, 80, 89},
			needle:   8,
			expected: 6,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			result := binary_search(tc.slice, tc.needle)

			// Check the result
			if result != tc.expected {
				t.Errorf("binary_search(%v, %d) = %d, expected %d", tc.slice, tc.needle, result, tc.expected)
			}
		})
	}
}
