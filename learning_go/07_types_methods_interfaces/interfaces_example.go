package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return "processed: " + data
	// program login
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program(data string) string {
	// get data
	return c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	fmt.Println(c.Program("example data"))
}
