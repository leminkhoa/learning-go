package main

import "fmt"

func main() {
	s := []string{"first", "second", "third"}
	fmt.Println(s)

	// Clear takes in a slice and sets all of the slice's elements to their zero value (for string, it is "")
	// Length of slice remains unchanged
	clear(s)
	fmt.Println(s, len(s))
}
