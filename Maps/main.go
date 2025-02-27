package main

import (
	"fmt"
	"maps" // This package is only available in Go 1.21+ for map comparison.
)

func main() {

	// Creating an empty map with string keys and int values
	m := make(map[string]int)

	// Assigning key-value pairs
	m["k1"] = 7
	m["k2"] = 13

	// Printing the entire map
	fmt.Println("map:", m)

	// Retrieving a value from the map using a key
	v1 := m["k1"]
	fmt.Println("v1:", v1) // Output: 7

	// Accessing a non-existent key returns the zero value (0 for int)
	v3 := m["k3"]
	fmt.Println("v3:", v3) // Output: 0

	// Getting the number of key-value pairs in the map
	fmt.Println("len:", len(m)) // Output: 2

	// Deleting a key from the map
	delete(m, "k2")
	fmt.Println("map:", m) // "k2" is removed

	// Clearing all elements from the map (Go 1.21+)
	clear(m)
	fmt.Println("map:", m) // Now an empty map

	// Checking if a key exists in the map
	_, prs := m["k2"]        // prs will be false because "k2" was deleted
	fmt.Println("prs:", prs) // Output: false

	// Declaring and initializing a map with values
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // Output: map[foo:1 bar:2]

	// Another map with the same key-value pairs
	n2 := map[string]int{"foo": 1, "bar": 2}

	// Using maps.Equal (Go 1.21+) to compare two maps
	if maps.Equal(n, n2) {
		fmt.Println("n == n2") // Output: n == n2
	}

	// Iterating over a map using a for loop
	for k, v := range n {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
