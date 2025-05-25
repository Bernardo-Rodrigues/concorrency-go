package semaphore

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{} // acquire semaphore (blocks if limit reached)
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	<-sem // release semaphore
}

func Execute() {
	const numWorkers = 5
	const maxConcurrent = 2

	sem := make(chan struct{}, maxConcurrent) // semaphore with limit of concurrent workers
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	wg.Wait()
}
