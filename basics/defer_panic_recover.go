package main

import "fmt"

func dpr() {

	// DEFER

	// normal control flow
	fmt.Println("first")
	fmt.Println("second")
	fmt.Println("third")

	// deferred execution
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

	// PANIC

	fmt.Println("start")
	panic("Panic move!")
	fmt.Println("end")

	// it's possible to re-panic function

	// RECOVER
	// can only be used inside deferred functions,
	// because deferred functions will be executed even after panic
}
