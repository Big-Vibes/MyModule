package main

import "fmt"

func main() {
	fmt.Println("Mathematics Score for victoria")
	score := 99
	if score > 75 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Fail")
	}

	grade := 6
	switch grade {
	case 1:
		fmt.Println("f")
	case 2:
		fmt.Println("e")
	case 3:
		fmt.Println("d")
	case 4:
		fmt.Println("c")
	case 5:
		fmt.Println("b")
	case 6:
		fmt.Println("A")
	}

	fmt.Println("Count 1 to 5")
	for i := 1; i <= 6; i++ {
		fmt.Println(i)
	}

}
