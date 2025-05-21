package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x) // First is destination and second is the source slice - return the number of elements copied
	fmt.Println(y, num)

	/** You can also copy a subset of a slice **/
	a := []int{1, 2, 3, 4}
	b := make([]int, 2)
	num_2 := copy(b, a)
	fmt.Println(b, num_2)

	/** Or copy from the middle of the source **/
	c := []int{1, 2, 3, 4}
	d := make([]int, 2)
	copy(d, c[2:])
	fmt.Println(d)
}
