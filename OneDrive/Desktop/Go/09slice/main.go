package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices is on")
	var fruitList = []string{"apple", "tomato", "pine"}
	fmt.Printf("type of fruitlis is: %T\n", fruitList)

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println("list is:", fruitList)
	fruitList = append(fruitList[1:3]) //last range is not included
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 324
	highScore[1] = 776
	highScore[2] = 566
	highScore[3] = 345
	//	highScore[4] = 588
	highScore = append(highScore, 555, 678, 345)
	fmt.Println(highScore)
	fmt.Println(len(highScore))
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	//	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove values from slices

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
