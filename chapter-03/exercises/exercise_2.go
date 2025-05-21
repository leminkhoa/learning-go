//go:build exercise2

package main

import "fmt"

// Write a program that defines a string variable called message with the value
// "Hi 👧 and 👦" and prints the fourth rune in it as a character, not a number.
func main() {
	message := "Hi 👧 and 👦"
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}
