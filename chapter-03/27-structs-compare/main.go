package main

import "fmt"

func main() {
	type firstPerson struct {
		name string
		age  int
	}

	// You can use type conversion to convert an instance of firstPerson to secondPerson
	// But you can't use == to compare
	type secondPerson struct {
		name string
		age  int
	}

	// You can't convert firstPerson to thirdPerson as fields are in different order
	type thirdPerson struct {
		age  int
		name string
	}

	// You can't convert firstPerson to fourthPerson because field names don't match
	type fourthPerson struct {
		firstName string
		age       int
	}

	// You can't convert firstPerson to fifthPerson because there is additional field
	type fifthPerson struct {
		name           string
		age            int
		favouriteColor string
	}

	// Anonymous struct can add a small twist
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g)

}
