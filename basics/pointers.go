package main

import "fmt"

func main() {

	// BASIC DEMO

	var a int = 42
	var b *int = &a   // b is a pointer of type int, which points to a
	fmt.Println(a, b)
	fmt.Println(&a, *b)   // & is referencing ("addres of") operator, * is dereferencing operator


	// POINTERS ARITHMETIC
	// a właściwie po prostu:
	// POINTERS
	// bo w go nie ma arytmetyki wskaźników w takim znaczeniu jak C czy Cpp

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Printf("%v %p %p\n", x, y, z)

	var ms *myStruct
	// check this also:
	// ms = &myStruct{foo: 42}
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}