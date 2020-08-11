package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning to Use Context Package")

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ctx, cancel = context.WithCancel(ctx)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Cancel the generation.")
		cancel()
	}()

	for n := range gen(ctx) {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("NUM := %v \n", n)
	}

	fmt.Printf("Program run for: %v", time.Now().Sub(start))
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dst)
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
