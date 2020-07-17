package main

import (
	"fmt"
	"math/rand"

	"wesionary.team/dipeshdulal/golang-test/handygenerators/repeat"
	"wesionary.team/dipeshdulal/golang-test/handygenerators/take"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	for num := range take.Take(done, repeat.Repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
	fmt.Println()

	randFn := func() interface{} {
		return rand.Int()
	}

	for num := range take.Take(done, repeat.Fn(done, randFn), 10) {
		fmt.Printf("%v \n", num)
	}

}
