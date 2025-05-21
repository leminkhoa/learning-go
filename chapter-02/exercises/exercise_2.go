package main

import "fmt"

// Write a program that declares a constant called value that can be assigned to
// both an integer and a floating-point variable. Assign it to an integer called i and a
// floating-point variable called f. Print out i and f.
func Exercise2() {
	fmt.Println("\nExercise 2:")
	const value = 10
	var i = value
	var f float32 = value

	fmt.Printf("i: %d, f: %g\n", i, f)

}
