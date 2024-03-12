package main

import "fmt"

func bckp_main() {
	ch := make(chan int)

	go func() {
		ch <- 42
		close(ch)
	}()

	value, ok := <-ch
	if ok {
		fmt.Println("Received value:", value)
	} else {
		fmt.Println("Channel closed")
	}

	value, ok = <-ch
	if ok {
		fmt.Println("Received value:", value)
	} else {
		fmt.Println("Channel closed")
	}
}
