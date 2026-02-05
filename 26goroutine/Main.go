package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signal = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatuscode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signal)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatuscode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoints")
		// panic(err)
	} else {
		mut.Lock()
		signal = append(signal, endpoint)
		mut.Unlock()

		fmt.Printf("%d Status code for %s\n", result.StatusCode, endpoint)
	}
}