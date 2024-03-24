package main

import (
	"context"
	"fmt"
)

type KeyName struct {
	name string
}

type KeyAge struct {
	age int
}

type KeyCity struct {
	city string
}

var (
	keyName = KeyName{name: "name"}
	keyAge  = KeyAge{age: 0}
	keyCity = KeyCity{city: "city"}
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, keyName, "John Doe")
	ctx = context.WithValue(ctx, keyAge, 30)
	ctx = context.WithValue(ctx, keyCity, "New York")
	fmt.Println(ctx.Value(keyName))
	fmt.Println(ctx.Value(keyAge))
	fmt.Println(ctx.Value(keyCity))
}
