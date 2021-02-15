package main

import (
	"bytes"
	"testing"

	"github.com/magiconair/properties/assert"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// There are two dependencies that we need to test
// 1) The number of calls to sleep
// 2) The order of the calls to sleep

type SuperSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *SuperSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SuperSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestHelloWorld(t *testing.T) {
	buffer := &bytes.Buffer{}

	t.Run("test the output buffer", func(t *testing.T) {

		superSpy := &SuperSpy{}
		Countdown(buffer, superSpy)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test the order of input", func(t *testing.T) {
		superSpy := &SuperSpy{}
		Countdown(superSpy, superSpy)
		got := superSpy.Calls
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		assert.Equal(t, got, want)
	})

	t.Run("test the number of calls", func(t *testing.T) {
		superSpy := &SuperSpy{}
		Countdown(superSpy, superSpy)
		var got int
		for _, item := range superSpy.Calls {
			if item == sleep {
				got++
			}
		}
		want := 4
		if got != want {
			t.Errorf("Wanted %v calls, got %v calls", want, got)
		}
	})
}
