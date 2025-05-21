package main

import "fmt"

func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)
}
