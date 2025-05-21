package main

import "fmt"

func main() {
	// You can use var to create a map variable that's set to its zero value
	var nilMap map[string]int // String key and int value
	fmt.Println(nilMap)

	// You can use := to create a map
	totalWins := map[string]int{}
	fmt.Println(totalWins)

	// Non empty map
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Raplh", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	// You can also use make to create a map with default size
	ages := make(map[int][]string, 10)
	fmt.Println(ages)

}
