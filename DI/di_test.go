package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// t.Fatal("not implemented")
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

}
