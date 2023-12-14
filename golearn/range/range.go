package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)
	// Set key/value pairs using typical name[key] = val syntax.

	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 132
	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.

	fmt.Println("map:", m)
	v := m["k1"]
	fmt.Println("v", v)

}
