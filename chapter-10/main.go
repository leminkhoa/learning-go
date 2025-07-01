package main

import (
	"fmt"

	/**
	`chapter-10` must be same as module name defined within go.mod
	**/
	math "chapter-10/02-package-math"
	format "chapter-10/03-package-format"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
