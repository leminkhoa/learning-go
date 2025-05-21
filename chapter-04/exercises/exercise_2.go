//go:build exercise_2

package main

import (
	"fmt"
	"math/rand"
)

/*
Loop over the slice you created in exercise 1. For each value in the slice, apply the
following rules:
a. If the value is divisible by 2, print “Two!”
b. If the value is divisible by 3, print “Three!”
c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
d. Otherwise, print “Never mind”.
*/
func main() {
	numbers := []int{}

	for i := 0; i <= 100; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	for _, v := range numbers {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")

		default:
			fmt.Println("Never mind")
		}
	}
}
