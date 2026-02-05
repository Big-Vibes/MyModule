package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("\nworlds ")

	fmt.Println("Hello")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
