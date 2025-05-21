package main

import "fmt"

func main() {
	// make allows you to specify the type, length, and, optionally, the capacity. Let’s take a look:
	x := make([]int, 5)
	fmt.Println(x)

	// You can create a slice with zero length but a capacity that's greater than zero
	y := make([]int, 0, 10)
	fmt.Println(y)

	// Since the length is 0, you can't directl index into it, but you can append value
	y = append(y, 5, 6, 7, 8)
	fmt.Println(y)

}
