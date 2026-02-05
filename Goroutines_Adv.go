package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing: %s\n", id, job)
	}
}

func main() {
	jobs := make(chan string, 5)
	var wg sync.WaitGroup

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Send jobs
	for _, j := range []string{"task1", "task2", "task3", "task4", "task5"} {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
