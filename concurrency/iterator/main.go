package main

import (
	"fmt"
	"time"
)

func fib(n int) chan int {
	c := make(chan int)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			time.Sleep(1 * time.Second)
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	fmt.Println("Iterator pattern.")
	fmt.Println("Generates fibonacci series every 1 second.")

	for i := range fib(200) {
		fmt.Printf("Got: %v \n", i)
	}

}
