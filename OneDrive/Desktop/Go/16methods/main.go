package main

import "fmt"

func main() {
	fmt.Println("learning methods")
	//no inheritance in golang : No super or parent class
	hitesh := User{"Ayush Roy", "royayush@mail.com", true, 23}
	//fmt.Println(hitesh)
	//fmt.Printf("ayush details are : %+v \n", hitesh)
	//fmt.Printf("Name is %v and email is %v .\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v .\n", hitesh.Name, hitesh.Email)
}

//this is a definition of struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@mail.com"
	fmt.Println("Email of this user is:", u.Email)
}
