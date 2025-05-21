package main

import "fmt"

func main() {
	var x []int
	x = append(x, 10)
	fmt.Println(x)

	// Append slice with multiple elements
	var y = []int{1, 2, 3}
	y = append(y, 5, 6, 7)
	fmt.Println(y)

	// Append one slice with anoother
	x = append(x, y...)
	fmt.Println(x)

}
