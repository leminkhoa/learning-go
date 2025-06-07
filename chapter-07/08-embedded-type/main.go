package main

import "fmt"

type Inner struct {
	X int
}
type Outer struct {
	Inner
	X int
}

func main() {
	// you can access the X on Inner only by specifying Inner explicitly:
	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)       // prints 20
	fmt.Println(o.Inner.X) // prints 10
}
