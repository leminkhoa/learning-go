//go:build exercise_3

package main

import "fmt"

func prefixer(prefix string) func(input string) string {
	return func(input string) string {
		return prefix + " " + input
	}
}

/*
*
Write a function called prefixer that has an input parameter of type string
and returns a function that has an input parameter of type string and returns a
string. The returned function should prefix its input with the string passed into
prefixer. Use the following main function to test prefixer
*
*/
func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
