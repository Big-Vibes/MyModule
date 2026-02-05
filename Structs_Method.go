package main

import "fmt"

	//Create a Stuct
	type User struct {
    Name  string
    Age   int
    Email string
}

//Method(Functions attached to structs)
func (u User) greet(){
	fmt.Println("Hello", u.Name)
	fmt.Println("Your Age", u.Age)
}

//Pointer Receiver
func (u *User) birthday(){
	u.Age++
}



func main() {
	user1 := User{
		Name:"Abiodun",
		Age: 24,
		Email: "abiodun@gmail.com",
}
	

	//access a struct
	fmt.Println(user1)
	fmt.Println(user1.Name)

	fmt.Println("\n")
	user1.greet()

	fmt.Println("\n")
	user1.birthday()
	fmt.Println("New Age: ",user1.Age)


}