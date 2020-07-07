package main

import (
	"fmt"
	"math/rand"
	"time"
)

func w1(c chan string) {
	for {
		r := rand.Int()
		c <- fmt.Sprintf("Worker1 -> %v ", r)
		time.Sleep(100 * time.Millisecond)
	}
}

func w2(c chan string) {
	for {
		r := rand.Int()
		c <- fmt.Sprintf("Worker2 -> %v ", r)
		time.Sleep(400 * time.Millisecond)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go w1(c1)
	go w2(c2)

	var x string
	// select to select between multiple channels
	for {
		select {
		case x = <-c1:
			fmt.Println(x)
		case x = <-c2:
			fmt.Println(x)
		}
	}
}
