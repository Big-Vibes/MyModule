package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointer")

	myNumber := 12
	var prt = &myNumber

	fmt.Println("Value of actual pointer ", *prt) //Print the actual value
	fmt.Println("Value of actual pointer ", prt)  //Print the actual memory location

	*prt = *prt * 2
	// *prt = *prt + 2
	fmt.Println(myNumber)
}
