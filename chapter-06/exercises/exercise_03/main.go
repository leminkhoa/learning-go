package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

/**
Write a program that builds a []Person with 10,000,000 entries (they could all
be the same names and ages). See how long it takes to run. Change the value
of GOGC and see how that affects the time it takes for the program to complete.
Set the environment variable GODEBUG=gctrace=1 to see when garbage collections
happen and see how changing GOGC changes the number of garbage collections.
What happens if you create the slice with a capacity of 10,000,000?
**/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	for _, gc := range []int{100, 10, 1000} {
		debug.SetGCPercent(gc)
		start := time.Now()
		data := []Person{}
		for i := 0; i < 10000000; i++ {
			data = append(data, Person{"khoa", "ab", 12})
		}
		elapsed := time.Since(start)
		fmt.Printf("GOGC=%d: %v\n", gc, elapsed)
	}
}
