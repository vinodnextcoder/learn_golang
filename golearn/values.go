package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // println add new line to the end
	fmt.Println("Vinod bhaua")
	// fmt.Printf("Hello go")

	// add number
	// A basic set of examples showing that %v is the default format, in this
	// case decimal for integers, which can be explicitly requested with %d;
	// the output is just what Println generates.
	integer := 23
	// Each of these prints "23" (without the quotes).
	fmt.Println(integer)

	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

}
