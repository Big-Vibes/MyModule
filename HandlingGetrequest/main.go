package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to get request examples")
	// ToPerformGetRequest()
	// PostJsonRequest()
	PerformPostFormRequest()
}

func ToPerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("content lenght: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is ", byteCount)
	fmt.Println(responseString.String())

	// Other ways
	// fmt.Println(string(content))
}

// To make Post

func PostJsonRequest() {
	const myurl = "http://localhost:8000/POST"

	// fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename":"Lets go with golang",
		"price":0,
		"platform":"cattyme.com"		
	}`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

// To Perform Post Form Request
func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstName", "Vibes")
	data.Add("LastName", "AI")
	data.Add("Email", "Vibes@go.com")
	data.Add("age", "8")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
