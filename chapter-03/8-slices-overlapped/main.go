package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] // a, b
	z := x[1:] // b, c, d
	x[1] = "y" // b -> y
	y[0] = "x" // a -> x
	z[1] = "z" // c -> z

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
