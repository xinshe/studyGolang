package main

import (
	"fmt"
	"time"
)

// worker pool （goroutine池）

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d: start job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d: end job %d\n", id, j)
		results <- j * 2
	}
}

func main()  {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	for j := 0; j < 5; j++ {
		jobs <- j
	}
	close(jobs)

	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 0; i < 5; i++ {
		<-results
	}

}
