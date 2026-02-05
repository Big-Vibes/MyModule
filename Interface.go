package main

import "fmt"

//Define an Interface
type speakerIN interface{
	Speak() string
}

//Implement Interface with Struct
type person struct{
	name string
}
func (p person) Speak() string{
	return "Hello, my name is " + p.name
}

//Implement Interface with Struct
type robot struct{
	id string
}
func (r robot) Speak() string{
	return "My RobotID is " + r.id
}

//implement a function that takes the interface as parameter
type Emailer struct{
	Email string
}
func (e Emailer) Speak() string{
	return "My Email is " + e.Email
}

//Use interface
func introduce (s speakerIN){
	fmt.Println(s.Speak())
}

func main() {
	Person1 := person{
		name: "Abiodun",
	}

	Robot1 := robot{
		id : "2001",
	}
	Emailer1 := Emailer{
		Email: "abioduninaolaji12@gmail",
	}

	

	introduce(Person1)
	introduce(Robot1)
	introduce(Emailer1)
} 