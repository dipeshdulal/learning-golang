package goroutines

import (
	"fmt"
	"time"
)

// Add2ValuesAfter2Seconds adds 2 values after 2 seconds
// and prints results accordingly
func Add2ValuesAfter2Seconds(a, b int) int {
	time.Sleep(2 * time.Second)
	fmt.Println("Value after addition is: ", a+b)
	return a + b
}
