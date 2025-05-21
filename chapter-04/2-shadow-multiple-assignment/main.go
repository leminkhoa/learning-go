package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

	/*
		Although there was an existing definition of x in an outer block, x was still shadowed
		within the if statement. That’s because := reuses only variables that are declared in
		the current block. When using :=, make sure that you don’t have any variables from
		an outer scope on the lefthand side unless you intend to shadow them.
	*/
}
