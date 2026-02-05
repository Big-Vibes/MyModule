package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go")

	var fruitlist [5]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[2] = "cerry"
	fruitlist[3] = "mango"
	fruitlist[4] = "okor"
	fmt.Println("This is your fruit list ", fruitlist)
	fmt.Println("This is your fruit list ", len(fruitlist))

	var vegList = [3]string{"potato", "beans", "room"}
	fmt.Println(vegList)

}
