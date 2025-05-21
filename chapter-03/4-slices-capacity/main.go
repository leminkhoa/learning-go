package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	var y = []int{1, 1, 1, 1, 1}
	x = append(x, y...)
	fmt.Println(x, len(x), cap(x))

	/**While it’s nice that slices grow automatically, it’s far more efficient to size them once.**/
}
