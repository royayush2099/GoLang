package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("making a web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type: %T\n", response)

	defer response.Body.Close() //callers responsibility to close the connection using this
	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)

}
