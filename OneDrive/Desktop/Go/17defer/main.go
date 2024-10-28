package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world") //this is going in lifo principal in reverser order first two then one and then world
	defer fmt.Println("one")
	defer fmt.Println("two") //this defer keyword puts this function at the last of main function
	fmt.Println("hellow")
	mydefer()
}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, "\n")
	}
}
