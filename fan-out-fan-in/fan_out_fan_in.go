package fan_out_fan_in

import (
	"fmt"
	"sync"
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Producer %d produced %d\n", id, i)
	}
}
func consumer(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in {
		out <- v * 2
		fmt.Printf("Consumer %d processed %d\n", id, v)
	}
}
func Execute() {
	numProducers := 2
	numConsumers := 2
	input := make(chan int, 10)
	output := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go producer(i, input, &wg)
	}
	wg.Wait()
	close(input)
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, input, output, &wg)
	}
	wg.Wait()
	close(output)
	for result := range output {
		fmt.Println("Result:", result)
	}
}
