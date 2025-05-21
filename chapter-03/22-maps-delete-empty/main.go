package main

import "fmt"

func main() {
	// Delete
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}

	fmt.Println(m)

	delete(m, "hello")
	fmt.Println(m)

	m = map[string]int{
		"hello": 5,
		"world": 10,
	}
	clear(m)
	fmt.Println(len(m))

}
