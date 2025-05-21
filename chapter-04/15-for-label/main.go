package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_n!"}

outer: // Used when there is nested loops
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}
