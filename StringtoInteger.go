// We need a function that can transform a string into a number. What ways of achieving this do you know?

// Note: Don't worry, all inputs will be strings, and every string is a perfectly valid representation of an integral number.

// Examples
// "1234" --> 1234
// "605"  --> 605
// "1405" --> 1405
// "-7" --> -7
package main

import (
	"strconv"
)

func StringToNumber(str string) (int, error) {
	// Use the Atoi function from the strconv package to convert the string to an integer.
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err // Return an error if the conversion fails.
	}
	return num, nil // Return the converted integer.
}

func main() {
	// Example usage:
	input := "1234"
	num, err := StringToNumber(input)
	if err != nil {
		// Handle the error here (e.g., invalid input).
	} else {
		// Use the converted number (num) here.
		println(num)
	}
}
