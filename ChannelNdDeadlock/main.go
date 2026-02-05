package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learning about channel in golang")

	myCh := make(chan int, 2) //size not to return error

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// Recieve Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, ischannelopen := <-myCh
		fmt.Println(ischannelopen)
		fmt.Println(val)

		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		close(myCh)
		// myCh <- 6

		// myCh <- 6

		wg.Done()
	}(myCh, wg)
	wg.Wait()

}
