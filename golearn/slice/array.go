package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// slice exampel
	myslice11 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice11)
	fmt.Printf("length = %d\n", len(myslice11))
	fmt.Printf("capacity = %d\n", cap(myslice11))

	// with omitted capacity
	myslice22 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice22)
	fmt.Printf("length = %d\n", len(myslice22))
	fmt.Printf("capacity = %d\n", cap(myslice22))
}
