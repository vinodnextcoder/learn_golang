package main

import "fmt"

func valutest() {
	var i int

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println("Your number is:", i)
	fmt.Println("Your number is:", i+1)
}
