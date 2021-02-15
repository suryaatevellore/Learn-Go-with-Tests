package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var finalWorld string = "Go!"
var constantStart int = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Tests the countdown from 3..2..1..Go
func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := constantStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(buffer, i)
	}
	sleeper.Sleep()
	fmt.Fprint(buffer, finalWorld)
}

func main() {
	defaultSleeper := &DefaultSleeper{}
	Countdown(os.Stdout, defaultSleeper)
}
