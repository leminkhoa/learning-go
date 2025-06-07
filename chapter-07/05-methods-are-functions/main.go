package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // prints 15

	// You can also assign the method to a variable or pass it to a parameter of type
	// func(int)int. This is called a method value:
	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	// You can also create a function from the type itself. This is called a method expression:
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 10))
}
