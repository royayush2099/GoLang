package main

import "fmt"

func main() {
	fmt.Println("welcome to a class of pointers")
	//var ptr *int
	//fmt.Println("value of ptr is:", ptr)
	mynum := 23
	var ptr *int = &mynum
	fmt.Println("value of ptr is:", ptr)
	fmt.Println("value of ptr is:", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:", mynum)
}
