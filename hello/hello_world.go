package main

import "fmt"

const englishHello = "Hello, "
const spanishHello = "Ola, "
const frenchHello = "Bonjour, "

// seperate your side effects from your actual code
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if language == "Spanish" {
		return spanishHello + name
	}

	if language == "French" {
		return frenchHello + name
	}
	return englishHello + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
