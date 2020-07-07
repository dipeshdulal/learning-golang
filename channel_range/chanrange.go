package main

import "fmt"

func fibo(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibo(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
