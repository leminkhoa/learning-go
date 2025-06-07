package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

// The value of iota increments for each constant in the const block, whether or not
// iota is used to define the value of a constant. The following code demonstrates what
// happens when iota is used intermittently in a const block:
const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

func main() {
	fmt.Println(Field1, Field2, Field3, Field4, Field5)
}

/**
Field2 is assigned 2 because iota has a value of 1 on the second line in the const
block. Field4 is assigned 20 because it has no type or value explicitly assigned, so it
gets the value of the previous line with a type and assignment. Finally, Field5 gets the
value 4 because it is the fifth line and iota starts counting from 0.
**/
