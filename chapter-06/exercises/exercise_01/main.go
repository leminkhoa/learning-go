package main

/**
Create a struct named Person with three fields: FirstName and LastName of
type string and Age of type int. Write a function called MakePerson that
takes in firstName, lastName, and age and returns a Person. Write a second
function MakePersonPointer that takes in firstName, lastName, and age and
returns a *Person. Call both from main. Compile your program with go build
-gcflags="-m". This both compiles your code and prints out which values
escape to the heap. Are you surprised about what escapes?
**/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	MakePerson("Khoa", "Le", 18)
	MakePersonPointer("Khoa", "Le", 18)
}
