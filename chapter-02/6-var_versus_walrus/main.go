package main

import "fmt"

func example_1() {
	// Most verbose way to declare a variable uses the var keyword
	var x int = 10

	// If the type on the righthand side of = is the expetected type, you can leave off the type from left side
	var y = 10

	// Conversely if you want to declare a variable and assign it the zero value
	var z int

	// You can also declare multiple vars
	var a, b int = 10, 20

	// You can declare all zero values of same type
	var c, d int

	// or different types
	var e, f = 10, "hello"

	fmt.Println(x, y, z, a, b, c, d, e, f)

}

func example_2() {
	// there are one more way to use var
	var (
		x    int
		y        = 20
		z    int = 30
		d, e     = 40, "hello"
		f, g string
	)

	fmt.Println(x, y, z, d, e, f, g)

}

func main() {
	example_1()
	example_2()
}
