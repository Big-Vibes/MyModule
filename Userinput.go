package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Enter your Rating"
	fmt.Println(Welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rate of pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)
	fmt.Printf("this is what type is %T", input)
}
