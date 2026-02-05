package main

import (
	// "fmt"
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Math in golang")

	// var num1 int = 2
	// var num2 float64 = 4

	// fmt.Println(num1 + int(num2))

	//Random Number in math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	//random in crypto
	MyranNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(MyranNum)

}
