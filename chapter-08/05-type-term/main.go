package main

import (
	"errors"
	"fmt"
)

// If you want a type term to be
// valid for any type that has the type term as its underlying type, put a ~ before the type
// term
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func divAndRemainder[T Integer](num, denom T) (T, T, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func main() {
	var a uint = 18_446_744_073_709_551_615
	var b uint = 9_223_372_036_854_775_808
	fmt.Println(divAndRemainder(a, b))

	type MyInt int // ~ allow to use user-defined types like this
	var myA MyInt = 10
	var myB MyInt = 20
	fmt.Println(divAndRemainder(myA, myB))

}
