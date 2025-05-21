//go:build exercise3

package main

import "fmt"

// Write a program that defines a struct called Employee with three fields:
// firstName, lastName, and id. The first two fields are of type string, and the
// last field (id) is of type int. Create three instances of this struct using whatever
// values you’d like. Initialize the first one using the struct literal style without
// names, the second using the struct literal style with names, and the third with a
// var declaration. Use dot notation to populate the fields in the third struct. Print
// out all three structs.
func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	employeeA := Employee{
		"Khoa",
		"Le",
		15,
	}
	fmt.Println(employeeA)

	employeeB := Employee{
		firstName: "Khoa",
		lastName:  "Le",
		id:        15,
	}
	fmt.Println(employeeB)

	var employeeC = Employee{}
	employeeC.firstName = "Khoa"
	employeeC.lastName = "Le"
	employeeC.id = 15
	fmt.Println(employeeC)
}
