package main

import "fmt"

func main() {
	randomNumbers := []int{5, 9, 12, 66, 3, 2, 10, 34, 60}

	// prepare common channel with inputs
	inputChan := generatePipeline(randomNumbers)

	// de-multiplex or fan-out to two channels
	c1 := squareNumber(inputChan)
	c2 := squareNumber(inputChan)

	// fan-in resulting numbers
	c := fanIn(c1, c2)

	sum := 0

	// Sum the data in combined channel
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}
	fmt.Printf("Total sum of squares: %d \n", sum)

}

func fanIn(c1, c2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
		}
	}()
	return c
}

func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}
