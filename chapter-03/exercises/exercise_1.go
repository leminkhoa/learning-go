//go:build exercise1

package main

import "fmt"

// Write a program that defines a variable named greetings of type slice of
// strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは",
// and "Привіт". Create a subslice containing the first two values; a second subslice
// with the second, third, and fourth values; and a third subslice with the fourth and
// fifth values. Print out all four slices.
func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[3:]
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(greetings)

}
