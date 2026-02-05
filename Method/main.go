package main

import "fmt"

func main() {
	fmt.Println("Struct in Golang")
	//no inheritance in golang; no super or parent

	Vibes := User{"Abiodun", "abiodun@go.dev", true, 17}
	tibes := User{"Odun", "", false, 24}
	fmt.Println(tibes)
	fmt.Println(Vibes)

	// fmt.Printf("Vibes details: %+v\n", Vibes)
	// fmt.Printf("Name is %v and Email is %v.\n ", Vibes.Name, Vibes.Age)
	// fmt.Printf("Name is %v and Email is %v.\n ", Vibes.Name, tibes.Age)

	tibes.GetStatus()
	Vibes.GetNewMail()
	fmt.Println(Vibes)

}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active", u.status)
}

func (u User) GetNewMail() {
	u.Email = "AbiodunInaIna12@go"
	fmt.Println("Your new email is ", u.Email)
}
