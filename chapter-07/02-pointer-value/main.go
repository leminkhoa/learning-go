package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() { /* Pointer receiver */
	c.total++
	c.lastUpdated = time.Now()
}
func (c Counter) String() string { /* Value receiver */
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}

/**
Pointer receiver methods can modify the original struct, because they have its address.
Value receiver methods work on a copy, so modifications are lost when the method ends.
**/
