package main

import (
	"fmt"
	"regex_engine/parser"
)

func main() {
	regex := "a(b|c)*d"
	ctx := parser.Parse(regex)
	fmt.Printf("Parsed regex: %+v\n", ctx)
}
