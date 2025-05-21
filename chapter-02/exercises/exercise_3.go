package main

import (
	"fmt"
	"math"
)

// Write a program with three variables, one named b of type byte, one named
// smallI of type int32, and one named bigI of type uint64. Assign each variable
// the maximum legal value for its type; then add 1 to each variable. Print out their
// values.
func Exercise3() {
	fmt.Println("\nExercise 3:")
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64
	fmt.Println("b:", b, "smallI:", smallI, "bigI:", bigI)

	// Add one to each variable
	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1
	fmt.Println("b:", b, "smallI:", smallI, "bigI:", bigI)

}
