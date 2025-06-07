package main

import "fmt"

type MyInt int

func typeAssert() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) /* variable i2 is of type MyInt */
	fmt.Println(i2 + 1)
}

func typeAssertPanicWrongType() {
	// we are using recover to allow us to run through the
	// failing type assertions. recover is explained in chapter 9.
	defer func() {
		if m := recover(); m != nil {
			fmt.Println(m) // prints out because a panic happens
		}
	}()
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(string)
	fmt.Println(i2)
}

func typeAssertPanicTypeNotIdentical() {
	// we are using recover to allow us to run through the
	// failing type assertions. recover is explained in chapter 9.
	defer func() {
		if m := recover(); m != nil {
			fmt.Println(m) // prints out because a panic happens
		}
	}()
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(int)
	fmt.Println(i2 + 1)
}

func typeAssertCommaOK() error {
	var i any
	var mine MyInt = 20
	i = mine
	i2, ok := i.(int)
	if !ok {
		// we are constructing a new error with fmt.Errorf.
		// fmt.Errorf is covered in chapter 9.
		return fmt.Errorf("unexpected type for %v", i)
	}
	fmt.Println(i2 + 1)
	return nil
}

func main() {
	typeAssert()
	typeAssertPanicWrongType()
	typeAssertPanicTypeNotIdentical()
	typeAssertCommaOK()
}
