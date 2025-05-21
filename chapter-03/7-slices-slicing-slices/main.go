package main

import "fmt"

func main() {
	// Slicing slides
	x := []string{"a", "b", "c", "d"}
	y := x[:2]  // Take until second element
	z := x[1:]  // Take after first element
	d := x[1:3] // Take after first element to third element
	e := x[:]   // Take all
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
