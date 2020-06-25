package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

func hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}
	return fmt.Sprintf("Hello %v!", user)
}
