package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Example on slices")

	var fruitlist = []string{"Apple", "Tomato", "peach"}
	// fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	//Another Example

	highScore := make([]int, 4)

	highScore[0] = 232
	highScore[1] = 503
	highScore[2] = 455
	highScore[3] = 999
	// highScore[4] = 1000  it wont run because it has a range of 4

	highScore = append(highScore, 500, 400, 300)

	fmt.Println(highScore)
	fmt.Println(len(highScore))

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//How to remove a value from slice on index

	var course = []string{"react", "js", "Swift", "python", "ruby"}

	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

}
