package main

import "testing"

func TestHello(t *testing.T) {

	assertgotwant := func(t *testing.T, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q , want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("world", "")
		want := "Hello, world"
		assertgotwant(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is applied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertgotwant(t, got, want)
	})

	t.Run("say 'Hello, world' in spanish", func(t *testing.T) {
		got := Hello("Surya", "Spanish")
		want := "Ola, Surya"
		assertgotwant(t, got, want)
	})

	t.Run("say 'Hello, world' in French", func(t *testing.T) {
		got := Hello("Surya", "French")
		want := "Bonjour, Surya"
		assertgotwant(t, got, want)
	})
}
