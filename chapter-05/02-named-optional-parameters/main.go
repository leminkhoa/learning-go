package main

import "fmt"

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)

	return nil
}

func main() {
	MyFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})

	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}
