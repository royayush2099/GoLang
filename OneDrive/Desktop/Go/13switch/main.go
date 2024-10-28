package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("learning switch case")
	rand.Seed(time.Now().UnixNano()) //it is depricated as of 1.20 version seed
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough //used to print 3 and the nex fallthrough need more on it
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and roll the dice again")
	}
}
