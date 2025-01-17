package select_with_timeout

import (
	"fmt"
	"time"
)

func Execute() {
	c := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c <- "result"
	}()

	select {
	case res := <-c:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
