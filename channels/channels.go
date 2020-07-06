package main

import "fmt"

func greet(c chan string) {
	<-c //Read from Channel
	<-c // Read Another from Channel
}

func main() {
	fmt.Println("main() started")

	c := make(chan string)

	go greet(c) // schedule a go routine

	c <- "John"

	close(c)

	c <- "Mike"

	fmt.Println("main() stopped")
}
