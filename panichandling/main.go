package main

import "fmt"

func main() {
	fmt.Println("Paninc Handling")
	printAllOperations(20, 23)

	// This should throw a panic
	printAllOperations(20, 0)
}

func printAllOperations(x, y int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from error: %v \n", r)
			fmt.Println("Proceeding to alternative flow.")
			printSkipDivide(x, y)
		}
	}()

	sum, subtract, multiply, divide := x+y, x-y, x*y, x/y
	fmt.Printf(
		"sum = %v, sub = %v, mul = %v, div = %v \n",
		sum,
		subtract,
		multiply,
		divide,
	)

}

func printSkipDivide(x, y int) {
	sum, subtract, multiply := x+y, x-y, x*y
	fmt.Printf(
		"sum = %v, sub = %v, mul = %v \n",
		sum,
		subtract,
		multiply,
	)
}
