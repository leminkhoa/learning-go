package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var fred person
	fmt.Println(fred)

	// By position
	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)

	// By key
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth)

	// A field in struct is accessed with dot notation
	beth.name = "Bethy"
	fmt.Println(beth)
}
