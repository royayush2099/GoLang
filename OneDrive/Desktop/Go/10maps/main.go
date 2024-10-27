package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["cpp"] = "c++"
	languages["go"] = "Go"

	fmt.Println("List of languages:", languages)
	fmt.Println("js shorts for:", languages["js"])
	delete(languages, "cpp") //cpp gone or deleted
	fmt.Println("List of languages:", languages)
	//loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("for key %v, value is %v \n", key, value)
	}
}
