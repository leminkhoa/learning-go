//go:build exercise_1

package main

import (
	"fmt"
	"math/rand"
)

/*
* Write a for loop that puts 100 random numbers between 0 and 100 into an int
slice
*
*/
func main() {
	numbers := []int{}

	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Intn(101))
	}

	fmt.Println(len(numbers), numbers)
}
