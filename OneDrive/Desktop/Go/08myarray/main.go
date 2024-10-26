package main

import "fmt"

func main() {
	fmt.Println("array is on ")
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "tomato"
	//fruitList[2] = " "
	fruitList[3] = "pine"
	fmt.Println("list is :", fruitList)
	fmt.Println("length is :", len(fruitList))
	var veglist = [3]string{"potato", "cabbage", "tomato"}
	fmt.Println("veg list is:", veglist)

}
