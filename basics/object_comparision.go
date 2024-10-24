package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user1 := User{"John", 25}
	user2 := User{"John", 25}

	if user1 == user2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
