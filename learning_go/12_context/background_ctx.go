package main

import (
	"context"
	"fmt"
)

func logic(ctx context.Context, s string) (string, error) {
	// Do some work
	return s, nil
}

func bckp_main() {
	ctx := context.Background()
	result, err := logic(ctx, "a string")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", result)
}
