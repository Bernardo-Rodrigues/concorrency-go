package fan_out_fan_in

import (
	"fmt"
	"sync"
)

// producer emits values into the input channel
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Producer %d produced %d\n", id, i)
	}
}

// consumer reads from input channel, processes the data, and sends it to the output channel
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

	// start producers
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go producer(i, input, &wg)
	}
	wg.Wait()    // wait for all producers to finish
	close(input) // no more values will be sent

	// start consumers
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, input, output, &wg)
	}
	wg.Wait()     // wait for all consumers to finish
	close(output) // no more values will be sent

	// collect results until channel is closed
	for result := range output {
		fmt.Println("Result:", result)
	}
}
