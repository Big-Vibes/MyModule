package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This need to go in a file, learn from vibes"
	file, err := os.Create("./Golangfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	CheckNilErr(err)

	length, err := io.WriteString(file, content)
	CheckNilErr(err)

	fmt.Println("lenght is: ", length)
	defer file.Close()

	Readfile("./Golangfile.txt")

}

func Readfile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	CheckNilErr(err)
	fmt.Println("The text data inside a file is \n", string(databyte))
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
