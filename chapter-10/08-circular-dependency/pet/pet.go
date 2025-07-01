package pet

import "github.com/learning-go-book-2e/ch10/sample_code/circular_dependency_example/person" /** This cause circular dependency **/

type Pet struct {
	Name      string
	Type      string
	OwnerName string
}

var owners = map[string]person.Person{
	"Bob":   {"Bob", 30, "Fluffy"},
	"Julia": {"Julia", 40, "Rex"},
}

func (p Pet) Owner() person.Person {
	return owners[p.OwnerName]
}
