package main

import "fmt"

// Sort ...
func Sort(s []int) {
	lowPtr := -1
	highPtr := len(s)
	for locPtr := 0; locPtr < highPtr; {
		if s[locPtr] == 0 {
			lowPtr++
			s[locPtr], s[lowPtr] = s[lowPtr], s[locPtr]
			locPtr++
		} else if s[locPtr] == 2 {
			fmt.Printf("Low: %v, High: %v; Loc: %v\n", lowPtr, highPtr, locPtr)
			highPtr--
			s[locPtr], s[highPtr] = s[highPtr], s[locPtr]
		} else {
			locPtr++
		}
	}
	fmt.Println(s)

}
func main() {
	s := []int{0, 1, 0, 1, 2, 2, 2, 0, 1}
	Sort(s)
}
