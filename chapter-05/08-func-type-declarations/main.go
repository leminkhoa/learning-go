package main

import "fmt"

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

type opFuncType func(int, int) int

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	fmt.Println(opMap)
}
