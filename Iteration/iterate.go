package main

// Iterate returns a string that contains a series of 'a' repeated n times
func Iterate(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}
