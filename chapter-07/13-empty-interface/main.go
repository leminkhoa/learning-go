package main

var i interface{}

func main() {
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}

	/*
		Because an empty interface doesn’t tell you anything about the value it represents, you
		can’t do a lot with it.
	*/
}
