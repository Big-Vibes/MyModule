package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")
	Grade := 70
	var points string

	if Grade >= 70 {
		points = "Excellent!!, You Got A"
	} else if Grade < 70 && Grade >= 50 {
		points = "Passed, You Got B "
	} else if Grade < 40 {
		fmt.Println("You failed the test")
	}
	fmt.Print(points)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err !=nil{

	// }

}
