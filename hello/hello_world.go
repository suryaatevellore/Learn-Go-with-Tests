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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHello
	case "French":
		prefix = frenchHello
	default:
		prefix = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
