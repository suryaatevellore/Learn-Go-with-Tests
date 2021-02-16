package main

import (
	"reflect"
	"testing"
)

func TestSorting(t *testing.T) {
	inputs := [][]int{
		{0, 0, 2, 2, 1, 1, 1},
		{0, 1, 0, 1, 0, 1, 0},
		{2, 1, 2, 1, 2, 1, 2},
	}

	expected := [][]int{
		{0, 0, 1, 1, 1, 2, 2},
		{0, 0, 0, 0, 1, 1, 1},
		{1, 1, 1, 2, 2, 2, 2},
	}

	for i := 0; i < 3; i++ {
		isEqual := reflect.DeepEqual(Sort(inputs[i]), expected[i])
		if !isEqual {
			t.Fatalf("Got %v, Expected %v", Sort(inputs[i]), expected[i])
		}

	}
}
