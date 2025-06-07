package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// something
	emp_a := Employee{"emp_a", "1"}
	emp_b := Employee{"emp_b", "2"}
	return []Employee{emp_a, emp_b}
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	// prints 12345
	fmt.Println(m.Description()) // prints Bob Bobson (12345)
}
