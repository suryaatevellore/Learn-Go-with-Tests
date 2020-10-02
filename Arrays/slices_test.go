package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	assertSum := func(t *testing.T, got, want int) {

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("test sum of elements of an array with varying sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertSum(t, got, want)
	})
}

func TestSumSlices(t *testing.T) {
	assertSumSlices := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("test sums of elements for arrays for varying slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{1, 2, 3})
		want := []int{3, 6}
		assertSumSlices(t, got, want)
	})
	t.Run("test sums of elements of array tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 1, 1, 1})
		want := []int{5, 3}
		assertSumSlices(t, got, want)
	})
	t.Run("test for an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		assertSumSlices(t, got, want)
	})

}
