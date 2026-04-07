package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxTimeoutExample()
}

func ctxTimeoutExample() {
	ctx := context.Background()
	ctxTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("called the API")
	case <-ctxTimeout.Done():
		fmt.Println("timeout: ", ctxTimeout.Err().Error())
	}
}
