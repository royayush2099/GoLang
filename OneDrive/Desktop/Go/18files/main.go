package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "this needs to go in a file"

	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("text data inside file is \n", string(databyte))
}
