package main

import "fmt"

// Adds two integers and returns their sum
func Adder(a, b int) int {
	return a + b
}
func main() {
	fmt.Println(Adder(2, 3))
}
