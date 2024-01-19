package main

import (
	"fmt"
	"time"
)

func comput(value int){

	for i:=0;i<value;i++{
		time.Sleep(time.Second)
		fmt.Println(i)
	}

}
func comput1(value int){

	for i:=0;i<value;i++{
		time.Sleep(time.Second)
		fmt.Println(i)
	}

}

func main() {
	comput(3)
	comput(4)
	fmt.Println("GO rountie")
	go comput(3)
	go comput(4)

}
