package main

import "fmt"

func main() {
	x := 10 // this assigment isn't read
	x = 20
	fmt.Println(x)
	x = 30 // this assigment isn't read

	/** While the compiler and go vet do not catch the unused assignments of 10 and 30 to
	x, third-party tools can detect them.**/
}
