package main

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	/* running this code shows that a function won’t
	change the values of the parameters passed into it
	*/
}
