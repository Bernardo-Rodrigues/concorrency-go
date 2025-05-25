package rate_limiting

import (
	"fmt"
	"time"
)

func Execute() {
	rate := time.Second
	ticker := time.NewTicker(rate) // limit processing to one request per second
	defer ticker.Stop()

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests) // no more incoming requests

	for req := range requests {
		<-ticker.C // wait for next tick before processing
		fmt.Println("Processing request", req)
	}
}
