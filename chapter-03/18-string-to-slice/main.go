package main

import "fmt"

func main() {
	var s string = "Hello, "
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs, cap(bs))
	fmt.Println(rs, cap(rs))
}
