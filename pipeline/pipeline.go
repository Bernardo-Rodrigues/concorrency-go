package pipeline

import "fmt"

// emit numbers to output channel
func stage1(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // signal end of data
	}()
	return out
}

// double each number
func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

// increment each number
func stage3(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1
		}
		close(out)
	}()
	return out
}

func Execute() {
	nums := []int{1, 2, 3, 4, 5}

	// pipeline: emit -> double -> increment
	c1 := stage1(nums) // channel emitting the input numbers
	c2 := stage2(c1)   // channel with doubled numbers
	c3 := stage3(c2)   // channel with incremented numbers

	for result := range c3 { // consume results until channel is closed
		fmt.Println(result)
	}
}
