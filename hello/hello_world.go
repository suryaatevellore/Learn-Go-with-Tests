package main

import "fmt"

const englishHello = "Hello, "

// seperate your side effects from your actual code
func Hello(name string) string {
	if name == "" {
		return englishHello + "world"
	}
	return englishHello + name
}

func main() {
	fmt.Println(Hello("world"))
}
