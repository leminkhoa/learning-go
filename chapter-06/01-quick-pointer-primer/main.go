package main

import "fmt"

func pointerExample() {
	// var x int32 = 10
	// var y bool = true
	// pointerX := &x
	// pointerY := &y
	// var pointerZ *string

	/**
	  | Variable  | Address | Value |
	  |-----------|---------|-------|
	  | x         | 1       | 0     |
	  |           | 2       | 0     |
	  |           | 3       | 0     |
	  |           | 4       | 10    |
	  | y         | 5       | 1     |
	  | pointerX  | 6       | 0     |
	  |           | 7       | 0     |
	  |           | 8       | 0     |
	  |           | 9       | 1     | --> x
	  | pointerY  | 10      | 0     |
	  |           | 11      | 0     |
	  |           | 12      | 0     |
	  |           | 13      | 5     | --> y
	  | pointerZ  | 14      | 0     | --> This pointer does not point to anything
	  |           | 15      | 0     |
	  |           | 16      | 0     |
	  |           | 17      | 0     |

	  **/
}

func pointerExample2() {

	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(&x)          // Same, also print a memory address of x
	fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	fmt.Println(z) // prints 15
}

func pointerExample3() {
	var x *int
	fmt.Println(x == nil) // print true
	fmt.Println(*x)       // panics
}

func main() {
	pointerExample()
	pointerExample2()
}
