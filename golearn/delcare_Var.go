package main

import (
	"fmt"
)

const PI = 3.14

func vartest() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Println(" Int varable")
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(" Const varable")
	const PI = 3.142
	fmt.Println(PI)
}
