package main

import "fmt"

func main() {
	// var a rune = 'x'
	// var s string = string(a)
	// var b byte = 'y'
	// var s2 string = string(b)

	/*A common bug for new Go developers is to try to make an int into
	a string by using a type conversion:
	*/
	var x int = 65
	var y = string(x) // This will be A instead of "65"
	fmt.Println(y)
}
