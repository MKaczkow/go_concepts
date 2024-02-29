// code from https://github.com/learning-go-book/pointer_performance
package main

import "fmt"

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFooNotGood(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 42
	return nil
}

func MakeFooGood() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 42,
	}
	return f, nil
}

func main() {
	f := Foo{}
	g := Foo{}
	MakeFooNotGood(&f)
	fmt.Println(f)
	g, _ = MakeFooGood()
	fmt.Println(g)
}
