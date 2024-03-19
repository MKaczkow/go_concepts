package main

import (
	"fmt"
	"sync"
)

func processAndGather(in <-chan int, processor func(int) int, num int) []int {
	// run monitoring goprocedure
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		// wait for all others
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processor(v)
			}
		}()
	}
	// call close on out channel
	go func() {
		wg.Wait()
		close(out)
	}()
	var result []int
	for v := range out {
		// for-range stops when out channel is closed
		result = append(result, v)
	}
	// return result
	return result
}

func main() {
	// prepare correct arguments for processAndGather func and call it
	input := make(chan int)
	processor := func(num int) int {
		return num * 2
	}

	go func() {
		defer close(input)
		for i := 0; i < 10; i++ {
			input <- i
		}
	}()

	result := processAndGather(input, processor, 10)

	fmt.Println(result)
}
