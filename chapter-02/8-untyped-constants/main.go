package main

func main() {
	// untyped constant
	const x = 10

	// following are legal
	var y int = x
	var z float64 = x
	var d byte = x

	// Here is what a typed constant declaration looks like
	const typedX int = 10
}
