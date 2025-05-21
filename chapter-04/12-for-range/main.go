package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// Ignore slice index in a for-range loop
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// What if you want the key but don't want the value
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}
