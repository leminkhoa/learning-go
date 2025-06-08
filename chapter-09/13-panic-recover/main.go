package main

import "fmt"

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}
}

/**
Running this code produces the following output:
60
30
runtime error: integer divide by zero
10
**/
