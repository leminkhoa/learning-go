package main

import "fmt"

/*
*
A shadowing variable is a variable that has the same name as a variable in a containing
block. For as long as the shadowing variable exists, you cannot access a shadowed
variable.
*
*/
func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
}
