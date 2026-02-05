package main

import "fmt"

func main() {

	//Create Maps
	age := map[string]int{
		"John" : 35,
		"Abiodun": 24,
		"Peter": 23,
	}

	//to Add a value
	age["Sarat"] = 40
	age["Faga"] = 100

	//to update a value
	age["John"] = 40

	//Delete a key
	delete (age, "Peter")

	//Check if Key Exist

	ages,Exist := age["Sarat"]
	if Exist{
		fmt.Println("Age of Sarat:", ages)
	} else {
		fmt.Println("peter doesn't Exist")
		fmt.Println("Sarat doesn't Exist")
	} 
	//Check if key Exist
	fmt.Println("\n")
	Checkage := "Abiodun"
	Checkages, existing := age[Checkage]
	if existing{
		fmt.Println(Checkage , "age is", Checkages)
	} else {
		fmt.Println("Abiodun name doesn't exist")
	}

	//Looping through a maps
	fmt.Println("\n")
	for name,ages := range age{
		fmt.Println(name,"is", ages, "years old")
	}
	//Access a value
	fmt.Println("\n")
	fmt.Println(age)
}