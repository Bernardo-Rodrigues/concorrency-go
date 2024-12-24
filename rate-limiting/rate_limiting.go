package rate_limiting

import (
	"fmt"
	"time"
)

func Execute() {
	rate := time.Second
	ticker := time.NewTicker(rate)
	defer ticker.Stop()

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for req := range requests {
		<-ticker.C // Esperar el siguiente tick
		fmt.Println("Processing request", req)
	}
}
