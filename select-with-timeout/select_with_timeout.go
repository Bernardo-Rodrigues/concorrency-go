package select_with_timeout

import (
	"fmt"
	"time"
)

func Execute() {
	c := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // simulate slow operation
		c <- "result"               // send result after delay
	}()

	select {
	case res := <-c:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second): // timeout if no response in 1s
		fmt.Println("Timeout")
	}
}
