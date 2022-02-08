package main

import (
	"fmt"
	"time"
)

func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Println("work id", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("work id", id, "finish job", j)
		result <- j * 2
	}
}

func main() {

	numJobs := 10
	job := make(chan int, numJobs)
	result := make(chan int, numJobs)

	// create job
	for i := 1; i <= 3; i++ {
		go worker(i, job, result)
	}

	// add job
	for j := 0; j <= numJobs; j++ {
		job <- j
	}

	close(job)

	for r := 0; r <= numJobs; r++ {
		fmt.Println("result is", <-result)
	}

}
