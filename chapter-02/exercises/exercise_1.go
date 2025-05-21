package main

import "fmt"

// Write a program that declares an integer variable called i with the value 20.
// Assign i to a floating-point variable named f. Print out i and f.
func Exercise1() {
	fmt.Println("Exercise 1:")
	var i int = 20
	var f float32 = float32(i)
	fmt.Printf("Integer: %d\nFloat: %g\n", i, f)
}
