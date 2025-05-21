package main

import "fmt"

func main() {
	/**
	The anonymous function assigned to f can read and write a, even though a is not
	passed in to the function.
	**/
	a := 20
	f := func() {
		fmt.Println(a)
		a = 30
	}
	f()
	fmt.Println(a)

	/**
	Using := instead of = inside the closure creates a new a that ceases to exist when
	the closure exits.
	**/
}
