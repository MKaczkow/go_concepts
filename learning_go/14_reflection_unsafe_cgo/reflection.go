package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := make([]int, 10)
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(reflect.ValueOf(v))

	someNumber := 42
	fmt.Println(reflect.TypeOf(someNumber))
	fmt.Println(reflect.ValueOf(someNumber))

	someString := "Hello, reflection!"
	someStringType := reflect.TypeOf(someString)
	someStringValue := reflect.ValueOf(someString)
	fmt.Println(someStringType)
	fmt.Println(someStringValue)
	fmt.Println(someStringType.Name())
	fmt.Println(someStringType.Kind())
}
