package main

import "fmt"

func main() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----")

	// or you can leave off the increment
	for i = 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}
