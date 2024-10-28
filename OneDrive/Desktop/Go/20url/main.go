package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://api.github.com:3000/users/royayush2099?type=User"

func main() {
	fmt.Println("handling urls")
	fmt.Println(myurl)
	//parsing
	result, _ := url.Parse(myurl)
	//	fmt.Println(result.Scheme)
	//	fmt.Println(result.Host)
	//	fmt.Println(result.Path)
	//fmt.Println(result.RawQuery)
	//fmt.Println(result.Port())
	qparams := result.Query()
	fmt.Printf("The types of  query params  are: %T\n", qparams)
	fmt.Println(qparams["type"])
}

//have to go deep
