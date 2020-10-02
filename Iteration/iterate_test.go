package main

import "testing"

func TestIterate(t *testing.T) {
	// t.Fatal("not implemented")
	got := Iterate("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("ca", 5)
	}
}

func ExampleIterate() {
	Iterate("a", 5)
}
