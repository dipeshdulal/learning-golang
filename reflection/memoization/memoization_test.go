package memoization_test

import (
	"fmt"
	"testing"
	"time"

	"wesionary.team/dipeshdulal/golang-test/reflection/memoization"
)

func TestCacherInput(t *testing.T) {
	mFnc := func(d int) int {
		return d * 2
	}
	memoization.Cacher(mFnc, 200)
	t.Logf("Success: takes input functions")

	d := 20
	if _, err := memoization.Cacher(d, 300); err != nil {
		t.Logf("Success: takes only input functions")
	} else {
		t.Errorf("Error: expected error but not thrown")
	}

	funcNoParam := func() {
		fmt.Println("just sample")
	}

	if _, err := memoization.Cacher(funcNoParam, 300); err != nil {
		t.Logf("Success: shouldnot take function with no parameters.")
	} else {
		t.Errorf("Error: expected error but not thrown.")
	}

	funcNoReturnValue := func(i int) {
		fmt.Printf("input is: %v \n", i)
	}

	if _, err := memoization.Cacher(funcNoReturnValue, 300); err != nil {
		t.Logf("Success: must have at-least one return value")
	} else {
		t.Errorf("Error: expected error must have one return value but not thrown")
	}

}

func TestActualMemoization(t *testing.T) {
	addSlowly := func(a, b int) int {
		fmt.Println("called function.")
		time.Sleep(100 * time.Millisecond)
		return a + b
	}

	ch, err := memoization.Cacher(addSlowly, 2*time.Second)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}

	chAddSlowly := ch.(func(int, int) int)

	for i := 0; i < 5; i++ {
		start := time.Now()
		result := chAddSlowly(1, 2)
		end := time.Now()
		fmt.Println("got result", result, "in", end.Sub(start))
	}

	// wait for cache to expire
	time.Sleep(3 * time.Second)
	start := time.Now()
	result := chAddSlowly(1, 2)
	end := time.Now()
	fmt.Println("got result ", result, "in", end.Sub(start))

}
