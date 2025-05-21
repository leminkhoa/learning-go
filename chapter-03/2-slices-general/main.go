package main

import (
	"fmt"
	"slices"
)

func main() {

	// Working with slices looks a lot like working with arrays, but subtle differences exist.
	//The first thing to notice is that you don’t specify the size of the slice when you declare
	//it:
	var x = []int{10, 20, 30} // Using [...] makes an array. Using [] makes a slice.
	fmt.Println(x)

	// You can also specify only the indices with nonzero values
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)

	// You can simulate multidimensional slices
	var z [][]int
	fmt.Println(z)

	// slice isn't comparable, the only thing you can compare a slice with using == is nil
	fmt.Println(z == nil)

	// Usng Go standard slices library
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	c := []int{1, 2, 3, 4, 5, 6}
	d := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(a, b)) // true
	fmt.Println(slices.Equal(a, c)) // false
	fmt.Println(slices.Equal(a, d)) // does not compile
}
