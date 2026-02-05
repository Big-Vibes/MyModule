package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://vibescatty.dev:3000/learn?coursename=reactjs&payment=2asf09fwef09"

func main() {
	fmt.Println("Url in Golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println("My url", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query parameters are %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("param is", val)
	}

	//to construct url from the value
	partOFurl := &url.URL{
		Scheme: "https",
		Host:   "127.0.0.1:5501",
		Path:   "/Practice/index.html",
	}
	anotherUrl := partOFurl.String()
	fmt.Println(anotherUrl)
}
