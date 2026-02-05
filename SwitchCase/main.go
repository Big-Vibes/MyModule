package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough //makes you print the next number below
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can open or move 6 spot")
	default:
		fmt.Println("roll the dice again")

	}
}
