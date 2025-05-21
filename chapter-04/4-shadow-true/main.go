package main

import "fmt"

func main() {
	fmt.Println(true)
	true := 10
	fmt.Println(true)

	// You must be very careful to never redefine any identifiers in the universe block
}
