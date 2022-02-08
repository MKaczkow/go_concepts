package main

import "fmt"

func main() {
	// normal control flow
	fmt.Println("first")
	fmt.Println("second")
	fmt.Println("third")

	//deferred execution
	// deffered statements are executed after main func,
	// but before it's return
	fmt.Println("fourth")
	defer fmt.Println("deferred fifth")
	fmt.Println("sixth")

	// as shown below
	// deferred statements are executed in LIFO order
	defer fmt.Println("deferred seventh")
	defer fmt.Println("deferred eighth")
	defer fmt.Println("deferred nineth")
}