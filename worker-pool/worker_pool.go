package worker_pool

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// exits cleanly when jobs channel is closed
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		// send result while results channel is still open
		results <- job * 2
	}
}

func Execute() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// prepare workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// close jobs channel to signal no more jobs will be sent
	close(jobs)

	wg.Wait()
	// close results channel after all workers are done
	close(results)

	// read all results from the closed channel
	for result := range results {
		fmt.Println("Result:", result)
	}
}
