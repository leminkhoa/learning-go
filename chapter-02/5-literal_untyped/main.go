package main

import "fmt"

func main() {
	// While you can’t add two integer variables together if they are declared to be of differ‐
	// ent types of integers, Go lets you use an integer literal in floating-point expressions or
	// even assign an integer literal to a floating-point variable:
	var x float64 = 10
	var y float64 = 200.3 * .5

	fmt.Println(x, y)

}
