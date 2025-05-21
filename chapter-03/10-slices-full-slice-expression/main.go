package main

import "fmt"

func main() {
	/**
	To avoid complicated slice situations, you should either never use append with a
	subslice or make sure that append doesn't cause an overwrite by using a full slice
	expression.

	The full slice expression includes a third
	part, which indicates the last position in the parent slice’s capacity that’s available
	for the subslice. Subtract the starting offset from this number to get the subslice’s
	capacity.
	**/
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")

	y := x[:2:2]
	z := x[2:4:4]

	fmt.Println(cap(x), cap(y), cap(z))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("after")
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
