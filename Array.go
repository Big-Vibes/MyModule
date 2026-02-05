package main

import "fmt"
	
func main(){
	// Declare an array
	var numbers [3]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	//Short declaration
	score := [4]int {1,2,3,4}
	
	//Create Slice
	Count := []int {1,2,5,10}

	//Slice a slice
	Count = append(Count,11)


	fmt.Println(numbers[0])
	fmt.Println(score[3])
	fmt.Println(Count[4])
	fmt.Println(Count[1:4])
	fmt.Println(Count[1] * Count[4])

	//Length and capacity
	fmt.Println(len(Count))
	fmt.Println(cap(Count))
	fmt.Println(cap(score))
	fmt.Println(len(score))
	fmt.Println(cap(numbers))
	fmt.Println(numbers)
	
}