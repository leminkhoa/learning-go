package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func main() {
	fmt.Println(divAndRemainder(12, 0))
}
