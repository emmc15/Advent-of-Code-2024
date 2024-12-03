package main

import (
	"testing"
)

type TestSet struct {
	expectedValue bool
	row           []int
}

func TestCheckRowIsValidWithTolerance(t *testing.T) {
	// test the function

	expectedSet := []TestSet{
		{true, []int{7, 6, 4, 2, 1}},
		{false, []int{1, 2, 7, 8, 9}},
		{false, []int{9, 7, 6, 2, 1}},
		{true, []int{1, 3, 2, 4, 5}},
		{true, []int{8, 6, 4, 4, 1}},
		{true, []int{1, 3, 6, 7, 9}},
		{true, []int{3, 3, 6, 7, 9}},
		{true, []int{1, 1, 2, 3, 4}},
	}

	for _, test := range expectedSet {
		result := checkRowIsValidWithTolerance(test.row, 1)
		if result != test.expectedValue {
			t.Errorf("Expected %v, but got %v", test.expectedValue, result)
		}
	}
}

func TestCheckRowIsValidBruteForce(t *testing.T) {
	// test the function

	expectedSet := []TestSet{
		{true, []int{7, 6, 4, 2, 1}},
		{false, []int{1, 2, 7, 8, 9}},
		{false, []int{9, 7, 6, 2, 1}},
		{true, []int{1, 3, 2, 4, 5}},
		{true, []int{8, 6, 4, 4, 1}},
		{true, []int{1, 3, 6, 7, 9}},
		{true, []int{3, 3, 6, 7, 9}},
		{true, []int{1, 1, 2, 3, 4}},
	}

	for index, test := range expectedSet {
		result := checkRowIsValidBruteForce(test.row)
		if result != test.expectedValue {
			t.Errorf("For Index %v, Expected %v, but got %v", index, test.expectedValue, result)
		}
	}
}
