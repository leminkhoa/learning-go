package person

import "github.com/learning-go-book-2e/ch10/sample_code/circular_dependency_example/pet" /** This cause circular dependency **/

type Person struct {
	Name    string
	Age     int
	PetName string
}

var pets = map[string]pet.Pet{
	"Fluffy": {"Fluffy", "Cat", "Bob"},
	"Rex":    {"Rex", "Dog", "Julia"},
}

func (p Person) Pet() pet.Pet {
	return pets[p.PetName]
}
