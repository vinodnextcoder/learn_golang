package main

import "fmt"

func zeroval(ival int) int {
	ival = 0
	return ival
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	var n = 0
	n = zeroval(i)
	fmt.Println("zeroval:", n)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
