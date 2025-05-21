package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt := "oops"
	fmt.Println(fmt) // (type string has no field or method Println)
}
