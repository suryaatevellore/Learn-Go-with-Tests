package main

import (
	"fmt"
	"testing"
)

func TestAddIntegers(t *testing.T) {
	// t.Fatal("not implemented")
	assertSum := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("test positive integers", func(t *testing.T) {
		got := Adder(2, 3)
		want := 5
		assertSum(t, got, want)
	})
}

func ExampleAdder() {
	sum := Adder(2, 3)
	fmt.Println(sum)
}
