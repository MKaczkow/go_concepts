package main

import (
	"fmt"
)

func failedUpdate(p *int) {
	i := 2137
	p = &i
}

func update(p *int) {
	*p = 2137
}

func bckp_main() {
	x := 10
	failedUpdate(&x) // won't change x
	fmt.Println(x)
	update(&x) // will change x
	fmt.Println(x)
}
