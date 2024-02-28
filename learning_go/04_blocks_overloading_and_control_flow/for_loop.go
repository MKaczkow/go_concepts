package main

import "fmt"

func bckp_main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
