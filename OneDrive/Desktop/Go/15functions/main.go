package main

import "fmt"

func main() {
	fmt.Println("welcome to function in golang")
	greeter()
	result := adder(3+4, 5+100)
	fmt.Println("Result is: ", result)
	proRes, myMessage := proAdder(8, 5, 6, 8, 9, 3, 45, 65)
	fmt.Println(proRes, myMessage)
}

//funciton definition is definitely going to outside in go
func greeter() {
	println("hello folks!")
}
func adder(valOne int, valTwo int) int { //this is function signature we are returning int value thats why
	return (valOne + valTwo)
}
func proAdder(values ...int) (int, string) { //we can give multiple args
	total := 0
	for _, val := range values {
		total += val

	}
	return total, "Hi pro result function" //sending strings with return values
}
