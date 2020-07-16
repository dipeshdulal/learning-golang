package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newRandStreamWithMemoryLeak() <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStream closure exited.")
		defer close(randStream)
		for {
			randStream <- rand.Int()
		}
	}()
	return randStream
}

func newRandStreamWithoutMemoryLeak(done <-chan interface{}) <-chan int {
	randStream := make(chan int)
	go func() {
		defer fmt.Println("newRandStreamWithoutMemoryLeak closure exited.")
		defer close(randStream)
		for {
			select {
			case randStream <- rand.Int():
			case <-done:
				return
			}
		}
	}()
	return randStream
}

func main() {
	randStream := newRandStreamWithMemoryLeak()
	fmt.Println("3 random ints.")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	done := make(chan interface{})
	randStreamWOLeak := newRandStreamWithoutMemoryLeak(done)
	fmt.Println("3 random ints.")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStreamWOLeak)
	}
	close(done)

	time.Sleep(1 * time.Second)
}
