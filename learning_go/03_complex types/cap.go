package main

import "fmt"

func bckp_main() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	y := []int{10, 20, 30, 40, 50}
	fmt.Println(y, len(y), cap(y))

	z := make([]int, 5, 10)
	fmt.Println(z, len(z), cap(z))

}
