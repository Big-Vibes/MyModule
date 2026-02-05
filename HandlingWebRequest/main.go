package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://127.0.0.1:5501/Practice/index.html"

func main() {
	fmt.Println("Vibes Catty Project")

	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("What is the response type %T\n", responce)

	defer responce.Body.Close() // caller's responsibility to close the response
	databytes, err := ioutil.ReadAll(responce.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
