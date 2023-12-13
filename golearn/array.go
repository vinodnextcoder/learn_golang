package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)
	car2 := [...]string{"dd", "aa"}
	fmt.Println(car2)
	// 	1:10 means: assign 10 to array index 1 (second element).
	// 2:40 means: assign 40 to array index 2 (third element).

	arr3 := [...]int{1: 10, 2: 40}

	fmt.Println(arr3)
	fmt.Println("Length of array 3", len(arr3))

}
