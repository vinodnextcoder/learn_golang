package main

import (
	"fmt"
)

func arraytest() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var cars = [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)
	car2 := [...]string{"dd", "aa"}
	fmt.Print(car2)

}
