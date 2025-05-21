package main

import "fmt"

func main() {
	// function is anonymous because they don't have a name
	f := func(j int) {
		fmt.Println("printing", j, "from inside of an anonymous function")
	}
	for i := 0; i < 5; i++ {
		f(i)
	}

}
