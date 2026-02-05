package main

import "fmt"

func main() {
	fmt.Println("Welcome to Function")
	greet()

	result := adder(4, 5, 4)
	fmt.Println("Result is", result)

	ProResult := Proadder(2, 3, 2, 3, 23, 23)
	// ProResult, myMessage := Proadder(2, 3, 2, 3, 23, 23)
	fmt.Println("ProResult is ", ProResult)
	// fmt.Println("my message", myMessage)
}

func adder(val1 int, val2 int, val3 int) int {
	return val1 + val2
}

func Proadder(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// func Proadder(vals ...int) (int, string) {
// 	total := 0
// 	for _, val := range vals {
// 		total += val
// 	}
// 	return total, "Harming"
// }

func greet() {
	fmt.Println("Ekaro to go")
}
