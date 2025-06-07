package main

import "fmt"

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func main() {
	doubledNumber, error := doubleEven(5)
	fmt.Println(doubledNumber)
	fmt.Println(error)

	doubledNumber, error = doubleEven(6)
	fmt.Println(doubledNumber)
	fmt.Println(error)
}
