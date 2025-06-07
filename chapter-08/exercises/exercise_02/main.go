/**
Define a generic interface called Printable that matches a type that implements
fmt.Stringer and has an underlying type of int or float64. Define types that
meet this interface. Write a function that takes in a Printable and prints its value
to the screen using fmt.Println.
**/

package main

import "fmt"

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type MyInt int

func (int MyInt) String() string {
	return fmt.Sprintf("%d", int)
}

type MyFloat float64

func (float MyFloat) String() string {
	return fmt.Sprintf("%f", float)
}

func Print[T Printable](value T) {
	fmt.Println(value)
}

func main() {
	Print(MyInt(12))
	Print(MyFloat(1.23))

}
