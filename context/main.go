package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context Example Start...")
	contextExample()
}

func contextExample() {
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second) // 2 seconds timeout
	//ctxWithTimeout, cancel := context.WithTimeout(ctx, 4*time.Second) // 4 seconds timeout
	defer cancel()

	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		println("task is complete")
	case <-ctxWithTimeout.Done():
		println("task timeout occurred")
	}
}
