package main

import "fmt"

func greet(ch chan string) {
	ch <- "Hello from Goroutine"
	ch <- "Goroutines are lightweight threads managed by Go runtime"
	close(ch)
}

func main() {

	ch := make(chan string, 2)
	go greet(ch)

	for message := range ch {
		fmt.Println(message)
	}
	//Or alternative way
	// message1 := <-ch
	// fmt.Println(message1)
	// message2 := <-ch
	// fmt.Println(message2)

}
