package pipeline

import "fmt"

func stage1(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

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
	c1 := stage1(nums)
	c2 := stage2(c1)
	c3 := stage3(c2)
	for result := range c3 {
		fmt.Println(result)
	}
}
