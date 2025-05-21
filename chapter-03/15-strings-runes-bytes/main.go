package main

import "fmt"

func main() {
	// You can extract a single value from an array or a slice
	var s string = "Hello there"
	var b byte = s[6]

	// The slice expression notation that you used with arrays and slices also works with strings
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(b, s2, s3, s4)

}
