package main

import (
	"fmt"
	"time"
)

func main() {
	// For example, you represent a duration of 2 hours and 30 minutes
	d := 2*time.Hour + 30*time.Minute // d is of type time.Duration
	fmt.Println(d)

	//
	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2023-03-13 00:00:00 +0000")
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))

}
