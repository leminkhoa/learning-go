package main

import (
	"fmt"
	"reflect"
)

func main() {
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	fmt.Println(reflect.TypeOf(xSlice))

	// You can also convert a subset of an array into slice
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	fmt.Println(y)
	fmt.Println(z)

}
