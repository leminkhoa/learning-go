package main

import "fmt"

func main() {

	// You can specify the size of the array and the type of the elements in the array
	var x [3]int
	fmt.Println(x)

	// You can specify them with an array literal
	x = [3]int{10, 20, 30}
	fmt.Println(x)

	// Sparse array
	var y = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)

	// You can use == and != to compare two arrays
	var a = [...]int{1, 2, 3}
	var b = [3]int{1, 2, 3}
	fmt.Println(a == b)

	// Go has only one-dimensional arrays, but you can simulate multidimensional arrays
	var c [2][3]int
	fmt.Println(c)

	// Arrays in Go are read and written using bracket syntax
	x[0] = 10
	fmt.Println(x[2])

	// Built in functino len takes in an array and return its length
	fmt.Println(len(x))
}
