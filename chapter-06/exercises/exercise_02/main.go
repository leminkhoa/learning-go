package main

import "fmt"

/**
Write two functions. The UpdateSlice function takes in a []string and a
string. It sets the last position in the passed-in slice to the passed-in string. At
the end of UpdateSlice, print the slice after making the change. The GrowSlice
function also takes in a []string and a string. It appends the string onto the
slice. At the end of GrowSlice, print the slice after making the change. Call these
functions from main. Print out the slice before each function is called and after
each function is called. Do you understand why some changes are visible in main
and why some changes are not?
**/

func UpdateSlice(s []string, str string) {
	if len(s) > 0 {
		s[len(s)-1] = str
	}
	fmt.Println("After UpdateSlice:", s)
}

func GrowSlice(s []string, str string) {
	s = append(s, str)
	fmt.Println("After GrowSlice:", s)
}

func main() {
	inputSlice := []string{"Hello", "World", "my", "name", "is"}
	fmt.Println(inputSlice)

	UpdateSlice(inputSlice, "z") //
	fmt.Println("In main after Update Slice", inputSlice)

	GrowSlice(inputSlice, "d")
	fmt.Println("In main after GrowSlice:", inputSlice)

}

/**
Below is the output:
	[Hello World my name is]
	After UpdateSlice: [Hello World my name z]
	In main after Update Slice [Hello World my name z]
	After GrowSlice: [Hello World my name z d]
	In main after GrowSlice: [Hello World my name z]

Explain:
- Update Slice: Slices in Go are reference types. When you pass a slice to a function,
you’re passing a small struct (pointer to the underlying array, length, and capacity).
--> Modify elements changes the underlying array

- Grow Slice: The append function may create a new underlying array if the original
slice doesn’t have enough capacity.
-> The new slice (s) inside the function points to this new array, but the original slice in main is unchanged.
-> The change is local to the function unless you return the new slice and assign it back in main.
**/
