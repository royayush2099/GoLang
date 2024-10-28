package main

import "fmt"

func main() {
	fmt.Println("learning if and else") //control flow
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out "
	} else {
		result = "exactly 10 login count"
	}
	fmt.Println(result)

}
